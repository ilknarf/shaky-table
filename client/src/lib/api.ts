export const createAccount = async (data: FormData): Promise<Response> => {
    return fetch("https://localhost:8080/create_account", {
        method: "POST",
        headers: {
            "Content-type": "application/x-www-form-urlencoded",
        },
        body: data,
    });
}
