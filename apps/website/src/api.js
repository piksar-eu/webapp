export const subscribe = async (email) => {
    const res = await fetch(`http://localhost:8080/api/easyconnect/subscribe`, {
        credentials: "include",
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            email
        })
    })

    if (!res.ok) {
        return Promise.reject("error")
    }

    return Promise.resolve(res)
}