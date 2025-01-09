<script>
    import { Link } from "svelte-routing";
    import { subscribe as subscribeApiCall } from "../api.js"

    let emailInput;
    let subscribed = false
    let subscriptionErr = false

    const subscribe = async (event) => {
        event.preventDefault();

        try {
            await subscribeApiCall(emailInput.value)
            subscribed = true
            subscriptionErr = false
        } catch (err) {
            subscriptionErr = true
            console.error(err)
        }
    }
</script>

<div class="newsletter">
    <div class="newsletter__head">Zapisz się do newslettera</div>
    { #if subscribed }
        <div class="newsletter__subscribed_msg">Dziękujemy za zapisanie do newslettera.</div>
    { :else }
        <form class="newsletter__input" on:submit={subscribe}>
            <input type="email" required bind:this={emailInput} placeholder="Twój email">
            <button type="submit" class="button">Zapisz się</button>
        </form>
        { #if subscriptionErr }
            <div class="newsletter__subscription_msg error">Podczas zapisu wystąpił błąd. Spróbuj ponownie.</div>
        { /if }
    { /if }
    <div class="newsletter__terms">Klikając „Zapisz się”, wyrażasz zgodę na przetwarzanie Twojego adresu e-mail w celu otrzymywania naszego newslettera zgodnie z&nbsp;<Link to="/polityka-prywatnosci">Polityką&nbsp;Prywatności</Link>.</div>
</div>

<style lang="scss">
    .newsletter {
        max-width: 480px;
        display: flex;
        flex-direction: column;

        &__head {
            padding: 1rem 0;
            font-size: 1.6rem;
            color: var(--clr-primary);
        }

        &__input {
            display: flex;
            
            input {
                padding: 0.6em 1.2em;
                font-size: 1em;
                flex-grow: 1;
            }

            button {
                margin-left: 8px;
            }
        }

        &__terms {
            font-size: 0.6rem;
            padding: 0.5rem 0;
            color: #767373;
        }

        &__subscribed_msg {
            color: #04461c;
            font-size: 1.2rem;
        }

        &__subscription_msg.error {
            color: #650349;
        }
    }
</style>