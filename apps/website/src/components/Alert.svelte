<script>
    import { alert } from '../store';

    let a;

    $: {
        if ($alert) {
            a = {
                type: $alert.type ?? "info",
                msg: $alert.msg ?? "",
            };

            alert.set(undefined);
        }
    }
</script>

{#if a}
    <div class="container">
        <div class="alert alert-{a?.type}">{a.msg}</div>
    </div>
{/if}

<style lang="scss">
    .alert {
        align-self: center;
        width: 100%;
        align-self: center;
        padding: 10px;
        border-width: 1px;
        border-style: solid;
        margin: 10px;
        color: #383d41;
        background-color: #e2e3e5;
        border-color: #d6d8db;
        border-radius: 5px;
    }

    .alert-success {
        color: #155724;
        background-color: #d4edda;
        border-color: #c3e6cb;
    }

    .alert-error {
        color: #721c24;
        background-color: #f8d7da;
        border-color: #f5c6cb;
    }
</style>