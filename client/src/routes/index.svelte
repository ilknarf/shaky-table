<script lang="ts">
    let message = "";

    const handleSubmit = (e: Event) => {
        e.preventDefault();
        const form = e.currentTarget as HTMLFormElement;
        const user = form.elements.namedItem("username") as HTMLInputElement;
        message = user.value;
    }
</script>

<h1>Welcome to SvelteKit</h1>
<form on:submit={handleSubmit}>
    <span>
        <label for="username">Username:</label>
        <input id="username" type="text">
    </span>
    <span>
        <label for="password">Password:</label>
        <input id="password" type="password">
    </span>
    <input id="submit" type="submit" value="Create Account" />
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
