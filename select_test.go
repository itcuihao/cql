package cql

import (
	"cql/utils"
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	s := NewSelect().
		From("user a")
	sel := s.String()
	fmt.Println(sel)

	s.Column("a.name", "a.age").
		Where(map[string]map[string]map[string]interface{}{
			"AND": map[string]map[string]interface{}{
				utils.Like: map[string]interface{}{
					"a.name": "Pert",
				},
				utils.In: map[string]interface{}{
					"a.age": []interface{}{24, 25},
				},

				utils.Eq: map[string]interface{}{
					"a.sex": 1,
				},
				utils.Gl: map[string]interface{}{
					"a.phone": 110,
				},
			},
		})
	sel = s.String()
	fmt.Println(sel)
	fmt.Println(s.Values)

	s.Limit(10)
	s.Offset(1)
	s.Column().Order("a.age desc").Group("a.name")
	sel = s.String()
	fmt.Println(sel)
	fmt.Println(s.Values)
}

// func Benchmark_select(t *testing.B) {
// 	t.ResetTimer()
// 	for i := 0; i < t.N; i++ {
// 		s := NewSelect().
// 			Column("a.name", "a.age").
// 			From("user a").
// 			Where(map[string]interface{}{
// 				"a.name": "Pert",
// 				"a.age":  24,
// 			})
// 		_ = s.String()
// 	}
// }
