import type { ICpf } from "$lib/types/cpf";

export async function validateCpf(cpfNumber: string) {
  console.log("cpf to validate:", cpfNumber )
  try {
    const response = await fetch(`http://localhost:5000/api/v1/cpfValidator`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        cpf_number: cpfNumber,
      }),
    });

    const validate = await response.json();
    console.log("API RESPONSE", validate);

    const validationResult = validate.response.toString();
    console.log("Validation Result:", validationResult);

    return validationResult;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}

export async function createeCpf(cpfData: ICpf) {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/savecpf`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: cpfData.name,
        city: cpfData.city,
        state: cpfData.state,
        cpfNumber: cpfData.cpfNumber,
      }),
    });

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR",response.json)
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}

export async function updateCpf(cpfToUpdate:string ,cpfData: ICpf) {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/cpf/${cpfToUpdate}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: cpfData.name,
        city: cpfData.city,
        state: cpfData.state,
        cpfNumber: cpfData.cpfNumber,
      }),
    });

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR",response.json)
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}

export async function deleteCpf(cpfToDelete:string) {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/cpf/${cpfToDelete}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      }
    });

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR",response.json)
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}
export async function addCpfToBlockList(cpfToAdd:string) {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/addblocklistcpf/${cpfToAdd}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      }
    });

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR",response.json)
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}
export async function reemoveCpfToBlockList(cpfToReemovee:string) {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/removeblocklistcpf/${cpfToReemovee}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      }
    });

    const validate = await response.status;
    if ((await response.status) != 200) {
      console.log("ERROR",response.json)
      return false;
    }
    location.reload();
    return true;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}

export async function InfosCpf(): Promise<ICpf[]> {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/cpfs`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const json = await response.json();
    const cpfReturn = json.response;

    const cpfs: ICpf[] = cpfReturn.map((item: any) => ({
      name: item.name,
      city: item.city,
      state: item.state,
      cpfNumber: item.cpfNumber,
    }));
    console.log("CPFs: ", cpfs);
    return cpfs;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}

export async function getBlockListCpfs(): Promise<ICpf[]> {
  try {
    const response = await fetch(`http://localhost:5000/api/v1/blocklistcpfs`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const json = await response.json();
    const cpfReturn = json.response;

    const cpfs: ICpf[] = cpfReturn.map((item: any) => ({
      name: item.name,
      city: item.city,
      state: item.state,
      cpfNumber: item.cpfNumber,
    }));
    console.log("CPFs: ", cpfs);
    return cpfs;
  } catch (error) {
    console.error("Error retrieving Cpf:", error);
    throw error;
  }
}

