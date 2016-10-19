package main

import "gopkg.in/macaron.v1"

func main() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())
<<<<<<< HEAD
    m.Get("/hello", func() string {
        return "Hello world!"
    })
m.Run()
}
=======
    m.Get("/", func() string {
        return "Hello world!"
    })

    m.Get("/reverse/:name",  func(ctx *macaron.Context) {
    ctx.Data["Name"] = ctx.Params(":name")
    //ctx.Data["Name"] = reverse(ctx.Params(":name"))
    ctx.HTML(200, "hello")
    })

    m.Run()
}

/*func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}*/
>>>>>>> origin/master
