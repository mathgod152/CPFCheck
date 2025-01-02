export async function  validateCnpj(cnpjNumber: string) {
    try {
        const response = await fetch(
            `http://localhost:5000/api/v1/cnpjValidator`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    cpf_number: cnpjNumber,
                }),
            }
        );
        var validate = await response.json();
        console.log(validate)
        return validate;
    } catch (error) {
        console.error("Error retrieving tarefa:", error);
        throw error;
    }
}
