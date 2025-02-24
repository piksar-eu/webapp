export const logout = async () => {
    const res = await _post('/api/auth/logout', {})
    
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }

    return Promise.resolve();
}

const _post = async (endpoint, data) => {
    return await fetch(`${import.meta.env.VITE_API_URL}${endpoint}`, {
        credentials: "include",
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
}