package decoder

import (
	"errors"
	"fmt"
	"io"
	"reflect"

	"gobook/ch12/practice12_7/parser"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode(v any) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Ptr {
		return errors.New("expected pointer")
	}
	ev := vv.Elem()

	node, err := parser.Parse(d.r)
	if err != nil {
		return err
	}

	n := newNode(node)
	switch ev.Kind() {
	case reflect.Struct:
		return d.decodeStruct(ev, n)
	case reflect.Map:
		return d.decodeMap(ev, n)
	case reflect.Slice, reflect.Array:
		return d.decodeSlice(ev, n)
	default:
		return d.decodeUnit(ev, n)
	}
}

func (d *Decoder) decodeStruct(v reflect.Value, n *node) error {
	ns, err := n.Struct()
	if err != nil {
		return err
	}
	for name, fieldNode := range ns.fields {
		fv := v.FieldByName(name)
		if fv.Kind() == reflect.Invalid {
			continue
		}
		err := d.decodeUnit(fv, fieldNode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Decoder) decodeMap(v reflect.Value, n *node) error {
	ns, err := n.Struct()
	if err != nil {
		return err
	}
	valueKind := v.Type().Elem().Kind()

	for name, fieldNode := range ns.fields {
		k, ok := d.findKey(v, name)
		if !ok {
			vv, err := d.getUnitValue(valueKind, fieldNode)
			if err != nil {
				return err
			}
			kv := reflect.ValueOf(name)
			v.SetMapIndex(kv, vv)
			continue
		}
		vv := v.MapIndex(k)
		if vv.Kind() == reflect.Invalid {
			continue
		}
		err := d.decodeUnit(vv, fieldNode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Decoder) findKey(v reflect.Value, key string) (reflect.Value, bool) {
	keys := v.MapKeys()
	for _, k := range keys {
		if k.Kind() != reflect.String {
			continue
		}
		if k.String() == key {
			return k, true
		}
	}

	return reflect.Value{}, false
}

func (d *Decoder) decodeSlice(v reflect.Value, n *node) error {
	ln, err := n.listNode()
	if err != nil {
		return err
	}
	vs := v.Len()
	var count int
	for _, child := range ln.Nodes() {
		if count >= vs {
			break
		}
		err := d.decodeUnit(v.Index(count), child)
		if err != nil {
			return err
		}
		count++
	}
	return nil
}

func (d *Decoder) decodeUnit(v reflect.Value, n *node) error {
	uv, err := d.getUnitValue(v.Kind(), n)
	if err != nil {
		return err
	}
	v.Set(uv)
	return nil
}

func (*Decoder) getUnitValue(k reflect.Kind, n *node) (reflect.Value, error) {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		number, err := n.Number()
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(number), nil
	case reflect.String:
		s, err := n.String()
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(s), nil
	default:
		return reflect.Value{}, fmt.Errorf("unsupported type %v", k)
	}
}
