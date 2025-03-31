export const logout = async () => {
    const res = await _post('/api/auth/logout', {})
    
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }

    return Promise.resolve();
}
export const getLeads = async () => {
    const res = await _get('/api/easyconnect/leads')
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    return res.json()
}

export const getUsers = async () => {
    const res = await _get('/api/auth/users')
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    return res.json()
}

export const getUser = async (id) => {
    const res = await _get(`/api/auth/users/${id}`)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    return res.json()
}

export const saveUser = async (id, name, roles) => {
    const res = await _post(`/api/auth/users/${id}`, {name, roles})
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
}

export const getRoles = async () => {
    const res = await _get('/api/auth/roles')
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    return res.json()
}

export const getRole = async (id) => {
    const res = await _get(`/api/auth/roles/${id}`)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    return res.json()
}

export const saveRole = async (id, name, permissions) => {
    const url = id ? `/api/auth/roles/${id}` : '/api/auth/roles'
    const res = await _post(url, {name, permissions})
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
}

export const deleteRole = async (id) => {
    const res = await _delete(`/api/auth/roles/${id}`)
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
}

export const getPermissions = async () => {
    const res = await _get('/api/auth/permissions')
    if (!res.ok) {
        return Promise.reject('err' in res ? res.err : 'err')
    }
    return res.json()
}

const _get = async (endpoint) => {
    if (import.meta.env.SSR) {
        return Promise.resolve(Response.json({}))
    }

    return await fetch(`${import.meta.env.VITE_API_URL}${endpoint}`, {
        credentials: "include",
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }
    })
}

const _post = async (endpoint, data) => {
    if (import.meta.env.SSR) {
        return Promise.resolve(Response.json())
    }

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

const _delete = async (endpoint) => {
    if (import.meta.env.SSR) {
        return Promise.resolve(Response.json())
    }

    return await fetch(`${import.meta.env.VITE_API_URL}${endpoint}`, {
        credentials: "include",
        method: 'DELETE',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }
    })
}