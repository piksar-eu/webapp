export const subscribe = async (email) => {
    const res = await fetch(`${import.meta.env.VITE_API_URL}/api/easyconnect/subscribe`, {
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