<script lang="ts">
    import { createAccount } from "$lib/api";
    
    let message = "";
    let alert = false;

    const handleSubmit = async (e: Event) => {
        e.preventDefault();
        const form = e.currentTarget as HTMLFormElement;
        // see https://github.com/microsoft/TypeScript/issues/30584
        // FormData might contain File, so coercing is needed :(
        const data = new URLSearchParams(new FormData(form) as URLSearchParams);

        message = "loading...";

        try {
            const response = await createAccount(data);
            const body = await response.json()

            if (!response.ok) {
                alert = true;
                message = body["message"] ?? "Error!";
            } else {
                window.location.href = "/login";
            }
        } catch(e) {
            alert = true;
            message = "Error!";
        }
    }
</script>

<h1>Shaky Table</h1>
<form on:submit={handleSubmit}>
    <span>
        <label for="username">Username:</label>
        <input name="username" type="text">
    </span>
    <span>
        <label for="password">Password:</label>
        <input name="password" type="password">
    </span>
    <input id="submit" type="submit" value="Create Account" />
    <p class={alert? "alert" : null}>{message}</p>
</form>

<style>
    span {
        display: flex;
        flex-direction: column;
    }
    input {
        max-width: 10em;
    }
    p.alert {
        color: red;
    }
</style>
