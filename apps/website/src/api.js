export const subscribe = async (email) => {
    const res = await _post('/api/easyconnect/subscribe', {
        email
    })

    if (!res.ok) {
        return Promise.reject("error")
    }

    return Promise.resolve(res)
}

export const register = async (email, salt, verifier) => {
    const res = await _post('/api/auth/register', {
        email, salt, verifier
    })

    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err');
    }

    return Promise.resolve();
}

export const loginInit = async (email) => {
    const res = await _post('/api/auth/login', {
        email
    })
    
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }

    return res.json()
}

export const loginSRP = async (M1, A) => {
    const res = await _post('/api/auth/login', { M1, A })
    
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }

    return res.json()
}

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