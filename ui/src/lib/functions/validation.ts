import type { IValidate } from "$lib/types/validateResponse";
import { validateCpf } from "./cpf";
import { validateCnpj} from "./cnpj";

export async function validateCpforCnpj(
    cpfCnpj: string,
    typeOfRequest: string
  ): Promise<IValidate> {
    const validationResponse: IValidate = {
      validateColor: "red",
      validationMessage: "Invalido",
    };

    if (!cpfCnpj || cpfCnpj.trim() === "") {
      return validationResponse;
    }

    const isValid = cpfCnpj.length > 10; 

    if (isValid) {
      let apiResponse = "false";

      if (typeOfRequest === "CPF") {
        apiResponse = await validateCpf(cpfCnpj); 
        console.log("API RESPONSE",apiResponse)
      } else if (typeOfRequest === "CNPJ") {
        apiResponse = await validateCnpj(cpfCnpj); 
      } else {
        return validationResponse; 
      }

      if (apiResponse == "true") {
        validationResponse.validateColor = "green";
        validationResponse.validationMessage = "Valido";
      }
    }

    return validationResponse; // Default invalid response
  }