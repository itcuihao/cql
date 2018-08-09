package utils

import (
	"fmt"
	"strings"
)

// GetWherev where 处理
func GetWherev(filter map[string]map[string]map[string]interface{}) (fw string, fv []interface{}) {
	fk := make([]string, 0, len(filter))
	fv = make([]interface{}, 0, len(filter))
	for ao, wmap := range filter {
		for op, vmap := range wmap {

			for k, v := range vmap {
				ops, vs := getOp(op, v)
				fk = append(fk, strings.Join([]string{k, ops}, Space))
				fv = append(fv, vs)
			}
		}
		if ao == Or {
			fw += "(" + strings.Join(fk, Space+ao+Space) + ")"
		}

		if ao == And {
			if len(fw) > 0 {
				fw += Space + And + Space + strings.Join(fk, Space+ao+Space)
			} else {
				fw += strings.Join(fk, Space+ao+Space)
			}
		}
	}
	return
}

func getOp(op string, v interface{}) (ops string, vs interface{}) {
	switch op {
	case In:
		vo, ok := v.([]interface{})
		fmt.Println(vo, ok, v)
		if !ok {
			return
		}
		ops, vs = buildIn(vo), vo
	case Like:
		ops, vs = strings.Join([]string{Like, Questmark}, Space), fmt.Sprintf("%%%v", v)
	default:
		ops, vs = strings.Join([]string{op, Questmark}, Space), v
	}
	return
}

func buildIn(vals []interface{}) (cond string) {
	cond = strings.TrimRight(strings.Repeat(Questmark+Comma, len(vals)), Comma)
	cond = fmt.Sprintf("%s (%s)", In, cond)
	return
}
