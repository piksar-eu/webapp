let _srpClient;

/**
 * @returns {Promise<ReturnType<typeof import('@swan-io/srp').createSRPClient>>}
 */
const srpClient = async () => {
    if (_srpClient == undefined) {
        const { createSRPClient } = await import('@swan-io/srp');
        _srpClient = createSRPClient("SHA-256", 3072);
    }

    return _srpClient
}

export { srpClient }