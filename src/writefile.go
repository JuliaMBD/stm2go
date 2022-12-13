package stm2go

// func Generate(domain string, pkgname string, xml []byte) {
// 	// default model names (to be modified)
// 	names := make(map[string]string)
// 	for i := 1; i < 100; i++ {
// 		names[strconv.Itoa(i)] = "stm" + strconv.Itoa(i)
// 	}

// 	stms, states := Parse(xml)
// 	pkg := NewGoPkgSource(domain, pkgname)
// 	stmap, sttree, root := NewGoSTMMap(pkg, names, stms, states)

// 	// generate directory
// 	if err := os.MkdirAll(pkgname+"/"+srcdir, 0777); err != nil {
// 		panic(err)
// 	}

// 	// generate common file
// 	if f, err := os.OpenFile(pkgname+"/"+srcdir+"/common.go", os.O_RDWR|os.O_CREATE, 0664); err == nil {
// 		w := NewWriter(f)
// 		pkg.Common(w)
// 		f.Close()
// 	} else {
// 		panic(err)
// 	}

// 	// generate base and impl files
// 	for _, s := range stmap {
// 		if f, err := os.OpenFile(pkgname+"/"+srcdir+"/"+s.Name+"_base.go", os.O_RDWR|os.O_CREATE, 0664); err == nil {
// 			w := NewWriter(f)
// 			s.BaseHeader(w)
// 			s.BaseStateDefinition(w)
// 			s.BaseStateInitialize(w)
// 			s.BaseTransDefinition(w)
// 			f.Close()
// 		} else {
// 			panic(err)
// 		}
// 		if f, err := os.OpenFile(pkgname+"/"+srcdir+"/"+s.Name+"_impl.go", os.O_RDWR|os.O_CREATE, 0664); err == nil {
// 			w := NewWriter(f)
// 			s.ImplHeader(w)
// 			s.ImplFunctions(w, sttree)
// 			f.Close()
// 		} else {
// 			panic(err)
// 		}
// 	}

// 	// generate test and main files
// 	if s, ok := sttree[root]; ok && len(s) == 1 {
// 		entryname := s[0].Name
// 		if f, err := os.OpenFile(pkgname+"/"+srcdir+"/"+pkgname+"_test.go", os.O_RDWR|os.O_CREATE, 0664); err == nil {
// 			w := NewWriter(f)
// 			pkg.TestGen(w, entryname)
// 			f.Close()
// 		} else {
// 			panic(err)
// 		}

// 		if f, err := os.OpenFile(pkgname+"/main.go", os.O_RDWR|os.O_CREATE, 0664); err == nil {
// 			w := NewWriter(f)
// 			pkg.GenMain(w, entryname)
// 			f.Close()
// 		} else {
// 			panic(err)
// 		}
// 	} else {
// 		fmt.Errorf("There are two or more root STMs.")
// 	}
// }
