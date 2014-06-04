package main

import (
	//"encoding/json"
	//"log"
	//"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoGonki", func() {

	// var (
	//     body []byte
	//     err error
	//     //todos []Todo
	// )

	Context("Get coords", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/coords", Coords)
			Expect(response.Code).To(Equal(200))
		})
	})
	Context("Get coords", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/coords", Coords)
			Expect(response.Code).To(Equal(200))
		})
	})

	Context("Get state", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/state", GetState)
			Expect(response.Code).To(Equal(200))
		})
	})

	Context("Move to new place", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/move/1/2", Move)
			Expect(response.Code).To(Equal(200))
		})
	})

	Context("Get profile data", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/profile", ShowProfile)
			Expect(response.Code).To(Equal(200))
		})
	})

	Context("User can login to system", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/login", Login)
			Expect(response.Code).To(Equal(200))
			Expect(response.Body).To(Equal("dashboard"))
		})
	})

	Context("User can logout from system", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/logout", Logout)
			Expect(response.Code).To(Equal(200))
		})
	})
	// Context("Create a Todo", func() {

	//     BeforeEach(func() {
	//         todo := Todo{"keep things green"}
	//         body, err = json.Marshal(todo)
	//         if err != nil {
	//             log.Println("Unable to marshal todo")
	//         }
	//     })

	//     It("returns a 200 Status Code", func() {
	//         PostRequest("POST", "/todos", HandleNewTodo, bytes.NewReader(body))
	//         Expect(response.Code).To(Equal(200))
	//     })
	// })

	// Context("Returns a single created todo", func() {
	//     It("returns a 200 Status Code", func() {
	//         Request("GET", "/todos", HandleIndex)
	//         log.Println(response)
	//         log.Println(response.Body)
	//         Expect(response.Code).To(Equal(200))
	//         Expect(response.Body).To(MatchJSON(`[{"Name":"keep things green"}]`))
	//     })
	// })
})
