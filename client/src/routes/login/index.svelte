<script lang="ts">
    import { loginUser } from "$lib/api";
    let message = "";

    const handleSubmit = async (e: Event) => {
        e.preventDefault();
        const form = e.currentTarget as HTMLFormElement;

        await loginUser(new FormData(form) as URLSearchParams);
    }
</script>

<h1>Shaky Table</h1>
<form on:submit={handleSubmit}>
    <span>
        <label for="username">Username:</label>
        <input id="username" type="text">
    </span>
    <span>
        <label for="password">Password:</label>
        <input id="password" type="password">
    </span>
    <input id="submit" type="submit" value="Login" />
    <p>{message}</p>
</form>

<style>
    span {
        display: flex;
        flex-direction: column;
    }
    input {
        max-width: 10em;
    }
    p {
        color: red;
    }
</style>
