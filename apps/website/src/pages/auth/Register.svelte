<script>
    import { register as registerApiCall } from '../../api';
    import { Link } from "svelte-routing";
    import { srpClient } from './srp-client';
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { alert } from '../../store';
    import Alert from '../../components/Alert.svelte';

    let emailInput, passwordInput;
    let isValid = false;

    const validate = (event) => {
        isValid =passwordInput.value && emailInput.reportValidity()
    }

     onMount(() => {
        emailInput?.focus();
    });

    const register = async (event) => {
        event.preventDefault();

        validate()
        if (!isValid) {
            return;
        }

        const client = await srpClient();

        const email = emailInput.value;
        const password = passwordInput.value;

        try {
            const salt = client.generateSalt();
            const privateKey = await client.derivePrivateKey(salt, email, password);
            const verifier = client.deriveVerifier(privateKey);

            await registerApiCall(email, salt, verifier)

            alert.set({type: "success", msg: "Zarejestrowano nowe konto! Możesz się zalogować."});

            navigate("/auth/login", { replace: true });
        } catch (e) {
            alert.set({type: "error", msg: "Błąd podczas rejestracji"});
        }
    }
</script>

<div class="centered-flex-layout">
    <Alert/>
    <form on:submit={register}>
        <input type="email" name="email" required placeholder="Twój email" bind:this={emailInput} on:input={validate}>
        <input type="password" name="password" required  placeholder="Hasło" bind:this={passwordInput} on:input={validate}>
        <button type="submit" class="button" disabled={!isValid}>Zarejestruj</button>
        <p> Masz już konto? <Link class="secondary" to="/auth/login">Zaloguj się!</Link></p>
    </form>
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

            p {
                margin-top: 20px;
            }
        }
    }
</style>