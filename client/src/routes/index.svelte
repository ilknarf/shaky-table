<script lang="ts">
    const handleSubmit = (e: Event) => {
        e.preventDefault();
    }
</script>

<h1>Welcome to SvelteKit</h1>
<form on:submit={handleSubmit}>
    <p>
        <label for="username">Username:</label>
        <input id="username" type="text">
    </p>
    <p>
        <label for="password">Password:</label>
        <input id="password" type="password">
    </p>
    <input id="submit" type="submit" value="Create Account" />
</form>

<style>
    p {
        display: flex;
        flex-direction: column;
    }
    input {
        max-width: 10em;
    }
</style>
