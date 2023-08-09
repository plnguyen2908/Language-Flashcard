const update = document.querySelector("#update")
update.disabled = true

const submit = document.querySelector("#submit")

const cancel = document.querySelector("#cancel")

const form = document.querySelector("form")

form.addEventListener("submit", e => {
    e.preventDefault()
    let collectionName = form.elements[0].value
    let word = form.elements[1].value
    let type = form.elements[2].value
    let definition = form.elements[3].value
    const register = async () => {
        try {
            const response = await fetch("http://localhost:8080/flash/" + collectionName, {
                method: "POST",
                headers: { "content-type": "application/json" },
                body: JSON.stringify({
                    word: word,
                    definition: [definition],
                    type: [type]
                })
            })
            if (response.status < 300) {
                const announcement = document.querySelector("#announcement")
                announcement.innerText = word + " has been added to " + collectionName + " successfully!"
            }
            else {
                const announcement = document.querySelector("#announcement")
                announcement.innerText = word + " has been added to " + collectionName + " already! If you want to add more definition to the word, please click the update button!"
                update.disabled = false
                submit.disabled = true
            }
        } catch (e) {
            const announcement = document.querySelector("#announcement")
            announcement.innerText = "Cannot add " + word + " to " + collectionName
            console.log(e)
        }
    }
    const signUp = async () => {
        try {
            const response = await fetch("http://localhost:8080/flash/" + collectionName, {
                method: "PUT",
                headers: { "content-type": "application/json" },
                body: JSON.stringify({
                    word: word,
                    definition: [definition],
                    type: [type]
                })
            })
            if (response.status < 300) {
                const announcement = document.querySelector("#announcement")
                announcement.innerText = word + " has been updated to " + collectionName + " successfully!"
                submit.disabled = false
                update.disabled = true
            }
            else {
                const announcement = document.querySelector("#announcement")
                announcement.innerText = "Cannot update " + word + " to " + collectionName
            }
        } catch (e) {
            const announcement = document.querySelector("#announcement")
            announcement.innerText = "Cannot update " + word + " to " + collectionName
            console.log(e)
        }
    }
    if (!submit.disabled)
        register()
    else
        signUp()
})

cancel.addEventListener("click", e => {
    submit.disabled = false
    update.disabled = true
    const announcement = document.querySelector("#announcement")
    announcement.innerText = ""
    document.querySelector("#set").value = ""
    document.querySelector("#word").value = ""
    document.querySelector("#type").value = ""
    document.querySelector("#definition").value = ""
})