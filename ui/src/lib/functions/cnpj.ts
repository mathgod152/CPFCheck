import type { ICnpj } from "$lib/types/cnpj";

export async function validateCnpj(cnpjNumber: string) {
  console.log("cnpj to validate:", cnpjNumber);
  try {
    const response = await fetch(`http://localhost:5000/api/v1/cnpjValidator`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        cnpj_number: cnpjNumber,
      }),
    });
    const validate = await response.json();
    console.log("API RESPONSE", validate);

    // Acessar e retornar a propriedade 'response' como string
    const validationResult = validate.response.toString();
    console.log("Validation Result:", validationResult);
    return validationResult;
  } catch (error) {
    console.error("Error retrieving tarefa:", error);
    throw error;
  }
}

export async function createeCnpj(cnpjData: ICnpj) {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/savecnpj`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: cnpjData.name,
        city: cnpjData.city,
        state: cnpjData.state,
        cnpjNumber: cnpjData.cnpjNumber,
      }),
    });

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR", response.json);
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}

export async function updateCnpj(cnpjToUpdate: string, cnpjData: ICnpj) {
  const encodedCnpj = cnpjToUpdate.replace(/\//g, "%2F");
  try {
    const response = await fetch(
      `http://localhost:5000/api/v1/cnpj/${encodedCnpj}`,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: cnpjData.name,
          city: cnpjData.city,
          state: cnpjData.state,
          cnpjNumber: cnpjData.cnpjNumber,
        }),
      }
    );

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR", response.json);
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}

export async function deleteCnpj(cnpjToDelete: string) {
  const encodedCnpj = cnpjToDelete.replace(/\//g, "%2F");
  try {
    const response = await fetch(
      `http://localhost:5000/api/v1/cnpj/${encodedCnpj}`,
      {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR", response.json);
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}
export async function addCnpjToBlockList(cnpjToAdd: string) {
  const encodedCnpj = cnpjToAdd.replace(/\//g, "%2F");
  try {
    const response = await fetch(
      `http://localhost:5000/api/v1/addblocklistcnpj/${encodedCnpj}`,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR", response.json);
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}
export async function reemoveCnpjToBlockList(cnpjToReemovee: string) {
  const encodedCnpj = cnpjToReemovee.replace(/\//g, "%2F");
  try {
    const response = await fetch(
      `http://localhost:5000/api/v1/removeblocklistcnpj/${encodedCnpj}`,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR", response.json);
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}

export async function InfosCnpj(): Promise<ICnpj[]> {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/cnpjs`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const json = await response.json();
    const cnpjReturn = json.response;

    const cnpjs: ICnpj[] = cnpjReturn.map((item: any) => ({
      name: item.name,
      city: item.city,
      state: item.state,
      cnpjNumber: item.cnpjNumber,
    }));
    console.log("Cnpjs: ", cnpjs);
    return cnpjs;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}

export async function getBlockListCnpjs(): Promise<ICnpj[]> {
  try {
    const response = await fetch(
      `http://localhost:5000/api/v1/blocklistcnpjs`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    const json = await response.json();
    const cnpjReturn = json.response;

    const cnpjs: ICnpj[] = cnpjReturn.map((item: any) => ({
      name: item.name,
      city: item.city,
      state: item.state,
      cnpjNumber: item.cnpjNumber,
    }));
    console.log("Cnpjs: ", cnpjs);
    return cnpjs;
  } catch (error) {
    console.error("Error retrieving Cnpj:", error);
    throw error;
  }
}
