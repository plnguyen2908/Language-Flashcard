let sets = []

const startingScreen = async () => {
    try {
        const resp = await fetch("http://localhost:8080/flash")
        const obj = await resp.json()
        sets = obj
        if (obj.length == 0) {
            const e = document.createElement("div")
            e.innerText = "There isn't any set. Please add one!"
            let id = document.querySelector("#header")
            id.children[1].append(e)
        }
        else {
            const e = document.createElement("div")
            e.innerText = "You currently have " + obj.length + " sets. Please type in the name to search!"
            let id = document.querySelector("#header")
            id.children[1].append(e)
        }
    } catch (e) {
        console.log("Error!", e)
    }
}

startingScreen()

const display = document.querySelector("#display")

display.addEventListener("click", e => {
    document.querySelector("#content").innerHTML = ""
    const getListname = async () => {
        try {
            const resp = await fetch("http://localhost:8080/flash")
            const obj = await resp.json()
            const ul = document.createElement("ul")
            for (let name of obj) {
                const li = document.createElement("li")
                li.innerText = name
                ul.append(li)
            }
            console.log(ul)
            document.querySelector("#content").append(ul)
        } catch (e) {
            console.log("Error!", e)
        }
    }
    getListname()
})

const reset = document.querySelector("#reset")

reset.addEventListener("click", e => {
    document.querySelector("#content").innerHTML = ""
    document.querySelector("input").value = ""
})

const createBlog = async (collectionName) => {
    const h1 = document.createElement("div")
    h1.innerHTML = "<h2>Presenting the words inside set " + collectionName + "</h2>"
    h1.style = "text-align: center;"
    content.append(h1)
    const resp = await fetch("http://localhost:8080/flash/" + collectionName)
    let obj = await resp.json()
    for (let element of obj) {
        const blog = document.createElement("div")
        blog.className = "blog"
        const word = document.createElement("div")
        word.innerText = "Word: " + element.word
        blog.append(word)
        for (let i = 0; i < element.definition.length; ++i) {
            const smallBlog = document.createElement("div")
            smallBlog.className = "smallBlog"
            const type = document.createElement("div")
            type.innerHTML = "Type: " + element.type[i]
            const definition = document.createElement("div")
            definition.innerHTML = "Definition: " + element.definition[i]
            smallBlog.append(type)
            smallBlog.append(definition)
            blog.append(smallBlog)
        }
        content.append(blog)
    }
}

const searching = (obj) => {
    const content = document.querySelector("#content")
    content.innerHTML = ""
    const collectionName = obj.value
    if (collectionName) {
        let ok = 0
        for (let element of sets)
            if (element === collectionName) {
                ok = 1
                break
            }
        if (ok) {
            createBlog(collectionName)
        }
        else {
            const h1 = document.createElement("div")
            h1.innerHTML = "<h2>The set " + collectionName + " does not exist</h2>"
            h1.style = "text-align: center;"
            content.append(h1)
        }
    }
}

const input = document.querySelector("input")

input.addEventListener("input", e => {
    searching(input)
})