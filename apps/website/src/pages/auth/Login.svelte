<script>
    import { Link } from "svelte-routing";
    import { loginInit as loginInitApiCall, loginSRP as loginSRPApiCall } from '../../api';
    import { srpClient } from './srp-client';
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
	import { alert, user } from '../../store';
    import Alert from "../../components/Alert.svelte";

	let emailInput;
	let passwordInput;
	let email = '';
    let salt;
    let B;
    let isEmailValid, isPasswordValid = false;

    onMount(() => {
        emailInput?.focus();
    });
    
    const validateEmail = () => {
        isEmailValid = emailInput?.reportValidity();
	}

	const validatePassword = () => {
        isPasswordValid = passwordInput?.reportValidity();
	}

    const loginInit = async (event) => {
        event.preventDefault();

        validateEmail()
        if (!isEmailValid) {
            return;
        }

        try {
            const r = await loginInitApiCall(emailInput.value ?? '')
            
            email = emailInput.value
            
            salt = r['SRP']['salt']
            B = r['SRP']['B']

        	setTimeout(() => passwordInput?.focus(), 10);
        } catch (e) {
            alert.set({type: "error", msg: "Błąd podczas logowania"});
        }
    }

    const loginSRP = async (event) => {
        event.preventDefault();
		
        validatePassword()
        if (!isPasswordValid) {
            return;
        }

        const client = await srpClient()

        const A = client.generateEphemeral();
        const privateKey = await client.derivePrivateKey(salt, email, passwordInput.value);
        const clientSession = await client.deriveSession(
            A.secret,
            B,
            salt,
            email,
            privateKey,
        );

        try {
            const r = await loginSRPApiCall(clientSession.proof, A.public);

            await client.verifySession(
                A.public,
                clientSession,
                r["M2"],
            );

            user.set(r["user"]);
            navigate("/", { replace: true });
        } catch (e) {
            alert.set({type: "error", msg: "Błąd podczas logowania"});
        }
    }
</script>

<div class="centered-flex-layout">
    <Alert/>
    {#if email == '' }
        <form on:submit={loginInit}>
            <input type="email" required bind:this={emailInput} placeholder="Twój email" on:input={validateEmail}>
            <button type="submit" class="button" disabled={!isEmailValid}>Dalej</button>

            <p> Nie masz jeszcze konta?  <Link class="secondary" to="/auth/register">Zarejestruj się!</Link></p>
        </form>
    {:else}
        <form on:submit={loginSRP}>
            <input type="email" readonly bind:value={email} placeholder="Twój email">
            <input type="password" required bind:this={passwordInput} placeholder="Hasło" minlength="4" on:input={validatePassword}>
            <button type="submit" class="button" disabled={!isPasswordValid}>Zaloguj się</button>

            <p> Nie masz jeszcze konta?  <Link class="secondary" to="/auth/register">Zarejestruj się!</Link></p>
        </form>
    {/if}
</div>

<style lang="scss">
    .centered-flex-layout {
        form {
            max-width: 300px;
            display: flex;
            flex-direction: column;

            input {
                margin-bottom: 5px;
            }

            input[readonly] {
                color: #a1a1a1;
            }

            p {
                margin-top: 20px;
            }
        }
    }
</style>