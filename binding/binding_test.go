package binding

import (
	"os"
	"path/filepath"
	"testing"
)

func TestError(t *testing.T) {
	_, err := NewServiceBinding()
	if err == nil {
		t.Error("expected error")
	}
}

func TestBindings(t *testing.T) {
	bd := t.TempDir()
	os.WriteFile(filepath.Join(bd, "junk"), []byte("junk text"), 0644)
	os.MkdirAll(filepath.Join(bd, "sb1", "sub1"), 0755)
	os.WriteFile(filepath.Join(bd, "sb1", "type"), []byte("mysql"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "username"), []byte("john"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "password"), []byte("L&ia6W@n7epi18a"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "url"), []byte("mysql://192.168.94.102:3306/school"), 0644)

	os.Mkdir(filepath.Join(bd, "sb2"), 0755)
	os.WriteFile(filepath.Join(bd, "sb2", "type"), []byte("neo4j"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "username"), []byte("jane"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "password"), []byte("o4%bGt#D8v2i0ja"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "url"), []byte("neo4j://192.168.94.103:7687/cr"), 0644)

	os.Setenv("SERVICE_BINDING_ROOT", bd)
	sb, err := NewServiceBinding()
	if err != nil {
		t.Error("expected no error")
	}

	l, err := sb.Bindings("mysql")
	if err != nil {
		t.Error("expected no error")
	}
	if len(l) != 1 {
		t.Fatal("expected length: 1, got: ", len(l))
	}

	v := l[0]
	if v["username"] != "john" {
		t.Fatal()
	}
	if v["password"] != "L&ia6W@n7epi18a" {
		t.Fatal()
	}
	if v["url"] != "mysql://192.168.94.102:3306/school" {
		t.Fatal()
	}

	l, err = sb.Bindings("neo4j")
	if err != nil {
		t.Error("expected no error")
	}
	if len(l) != 1 {
		t.Fatal("expected length: 1, got: ", len(l))
	}

	v = l[0]
	if v["username"] != "jane" {
		t.Fatal()
	}
	if v["password"] != "o4%bGt#D8v2i0ja" {
		t.Fatal()
	}
	if v["url"] != "neo4j://192.168.94.103:7687/cr" {
		t.Fatal()
	}

}

func TestBindingsWithProvider(t *testing.T) {
	bd := t.TempDir()
	os.WriteFile(filepath.Join(bd, "junk"), []byte("junk text"), 0644)
	os.MkdirAll(filepath.Join(bd, "sb1", "sub1"), 0755)
	os.WriteFile(filepath.Join(bd, "sb1", "type"), []byte("mysql"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "provider"), []byte("oracle"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "username"), []byte("john"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "password"), []byte("L&ia6W@n7epi18a"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "url"), []byte("mysql://192.168.94.102:3306/school"), 0644)

	os.Mkdir(filepath.Join(bd, "sb2"), 0755)
	os.WriteFile(filepath.Join(bd, "sb2", "type"), []byte("mysql"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "provider"), []byte("mariadb"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "username"), []byte("jane"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "password"), []byte("o4%bGt#D8v2i0ja"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "url"), []byte("mysql://192.168.94.103:7687/school"), 0644)

	os.Setenv("SERVICE_BINDING_ROOT", bd)
	sb, err := NewServiceBinding()
	if err != nil {
		t.Error("expected no error")
	}

	l, err := sb.BindingsWithProvider("mysql", "oracle")
	if err != nil {
		t.Error("expected no error")
	}
	if len(l) != 1 {
		t.Fatal("expected length: 1, got: ", len(l))
	}

	v := l[0]
	if v["type"] != "mysql" {
		t.Fatal()
	}
	if v["provider"] != "oracle" {
		t.Fatal()
	}
	if v["username"] != "john" {
		t.Fatal()
	}
	if v["password"] != "L&ia6W@n7epi18a" {
		t.Fatal()
	}
	if v["url"] != "mysql://192.168.94.102:3306/school" {
		t.Fatal()
	}

	l, err = sb.BindingsWithProvider("mysql", "mariadb")
	if err != nil {
		t.Error("expected no error")
	}
	if len(l) != 1 {
		t.Error("expected length: 1, got: ", len(l))
	}

	v = l[0]
	if v["type"] != "mysql" {
		t.Fatal()
	}
	if v["provider"] != "mariadb" {
		t.Fatal()
	}
	if v["username"] != "jane" {
		t.Fatal()
	}
	if v["password"] != "o4%bGt#D8v2i0ja" {
		t.Fatal()
	}
	if v["url"] != "mysql://192.168.94.103:7687/school" {
		t.Fatal()
	}

}

func TestAllBindings(t *testing.T) {
	bd := t.TempDir()
	os.WriteFile(filepath.Join(bd, "junk"), []byte("junk text"), 0644)
	os.MkdirAll(filepath.Join(bd, "sb1", "sub1"), 0755)
	os.WriteFile(filepath.Join(bd, "sb1", "type"), []byte("mysql"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "username"), []byte("john"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "password"), []byte("L&ia6W@n7epi18a"), 0644)
	os.WriteFile(filepath.Join(bd, "sb1", "url"), []byte("mysql://192.168.94.102:3306/school"), 0644)

	os.Mkdir(filepath.Join(bd, "sb2"), 0755)
	os.WriteFile(filepath.Join(bd, "sb2", "type"), []byte("neo4j"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "username"), []byte("jane"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "password"), []byte("o4%bGt#D8v2i0ja"), 0644)
	os.WriteFile(filepath.Join(bd, "sb2", "url"), []byte("neo4j://192.168.94.103:7687/cr"), 0644)

	os.Setenv("SERVICE_BINDING_ROOT", bd)
	sb, err := NewServiceBinding()
	if err != nil {
		t.Error("expected no error")
	}

	l, err := sb.AllBindings()
	if err != nil {
		t.Error("expected no error")
	}
	if len(l) != 2 {
		t.Error("expected length: 2, got: ", len(l))
	}

	mysqlFound := false
	neo4jFound := false

	for _, v := range l {
		if v["type"] == "mysql" {
			mysqlFound = true
			if v["username"] != "john" {
				t.Fail()
			}
			if v["password"] != "L&ia6W@n7epi18a" {
				t.Fail()
			}
			if v["url"] != "mysql://192.168.94.102:3306/school" {
				t.Fail()
			}
		}
		if v["type"] == "neo4j" {
			neo4jFound = true
			if v["username"] != "jane" {
				t.Fail()
			}
			if v["password"] != "o4%bGt#D8v2i0ja" {
				t.Fail()
			}
			if v["url"] != "neo4j://192.168.94.103:7687/cr" {
				t.Fail()
			}
		}
	}
	if !mysqlFound {
		t.Fail()
	}
	if !neo4jFound {
		t.Fail()
	}
}
