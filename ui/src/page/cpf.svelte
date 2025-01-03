<script lang="ts">
    import { onMount } from "svelte";
    import { validateCpforCnpj } from "$lib/functions/validation";
    import type { IValidate } from "$lib/types/validateResponse";
    import {
      createeCpf,
      deleteCpf,
      InfosCpf,
      updateCpf,
    } from "$lib/functions/cpf";
    import type { ICpf } from "$lib/types/cpf";
  
    // Variáveis de estado
    let validationMessage = "";
    let validationColor = "";
    let showValidationModal = false;
    let showCreationModal = false;
    let showUpdateModal = false;
    let showDeleteModal = false;
    let newCpfCnpj: ICpf = {
      name: "",
      city: "",
      state: "",
      type: "CPF",
      cpfNumber: "",
    };
    let validateCpfCnjpj = "";
    let currentType = "CPF";
    let data: ICpf[] = [];
    let filterState = "";
    let filterCity = "";
    let blocklist = [];
    let sortBy = "";
    let selectedItem: ICpf = {
      name: "",
      city: "",
      state: "",
      type: "CPF",
      cpfNumber: "",
    };
    let groupedCpf = false;
    $: filteredData();
  
    onMount(async () => {
      const loadedData = await InfosCpf(); // Carrega os dados
      data = [...loadedData]; // Cria uma nova referência, forçando a reatividade
      console.log("Dados carregados:", data);
    });
  
    async function handleCreateCpf() {
      const isCreated = await createeCpf(newCpfCnpj);
      if (isCreated) {
        console.log("CPF criado com sucesso!");
        showCreationModal = false; // Fecha o modal após o sucesso
        data = await InfosCpf(); // Atualiza os dados após criar
      } else {
        console.error("Falha ao criar o CPF.");
      }
    }
    async function handleUpdateCpf() {
      if (!selectedItem || !selectedItem.cpfNumber) {
        console.error("Nenhum item selecionado para atualizar.");
        return;
      }
  
      const isUpdated = await updateCpf(selectedItem.cpfNumber, selectedItem);
      if (isUpdated) {
        console.log("CPF atualizado com sucesso!");
        showUpdateModal = false; // Fecha o modal após o sucesso
      } else {
        console.error("Falha ao atualizar o CPF.");
      }
    }
    async function handleDeleteCpf() {
      if (!selectedItem || !selectedItem.cpfNumber) {
        console.error("Nenhum item selecionado para Deletar.");
        return;
      }
  
      const isDeleteed = await deleteCpf(selectedItem.cpfNumber);
      if (isDeleteed) {
        console.log("CPF atualizado com sucesso!");
        showUpdateModal = false; // Fecha o modal após o sucesso
      } else {
        console.error("Falha ao atualizar o CPF.");
      }
    }
  
    function filteredData() {
      let filtered = [...data]; // Cria uma cópia dos dados para filtrar
  
      // Aplicando o filtro por estado e cidade, se definidos
      if (filterState) {
        filtered = filtered.filter(
          (item) => item.state.trim() === filterState.trim()
        );
      }
  
      if (filterCity) {
        filtered = filtered.filter(
          (item) => item.city.trim() === filterCity.trim()
        );
      }
  
      // Ordenação por 'sortBy'
      if (sortBy && sortBy !== "") {
        filtered.sort((a, b) => (a[sortBy] > b[sortBy] ? 1 : -1));
      }
  
      return filtered;
    }
  
    // Funções do modal
    function openUpdateModal(item: ICpf) {
      selectedItem = { ...item };
      showUpdateModal = true;
    }
    function openDeleteModal(item: ICpf) {
      selectedItem = { ...item };
      showDeleteModal = true;
    }
  
    function closeUpdateModal() {
      selectedItem = null;
      showUpdateModal = false;
    }
    function closeDeleteModal() {
      selectedItem = null;
      showDeleteModal = false;
    }
  
    function toggleGroupByCpf() {
      groupedCpf = !groupedCpf;
    }
  
    $: groupedData = groupedCpf
      ? Object.values(
          data.reduce((acc, item) => {
            if (!acc[item.cpfNumber]) {
              acc[item.cpfNumber] = { ...item, count: 0 };
            }
            acc[item.cpfNumber].count++;
            return acc;
          }, {})
        )
      : data;
  
    // Validação
    async function validateInput() {
      const validationResponse: IValidate = await validateCpforCnpj(
        validateCpfCnjpj,
        currentType
      );
      validationMessage = validationResponse.validationMessage;
      validationColor = validationResponse.validateColor;
    }
  
    function openValidationModal() {
      showValidationModal = true;
    }
  
    function closeValidationModal() {
      showValidationModal = false;
      validationMessage = "";
      validationColor = "";
    }
  
    function openCreationModal() {
      showCreationModal = true;
    }
  
    function closeCreationModal() {
      showCreationModal = false;
      validateCpfCnjpj = "";
    }
  </script>
    {#if data.length > 0}
      <div class="table">
        <div class="row header-row">
          <div class="cell">Name</div>
          <div class="cell">CPF/CNPJ</div>
          <div class="cell">City</div>
          <div class="cell">State</div>
          <div class="cell">Actions</div>
        </div>
  
        {#each filteredData() as item}
          <div class="row">
            <div class="cell">{item.name}</div>
            <div class="cell">{item.cpfNumber}</div>
            <div class="cell">{item.city}</div>
            <div class="cell">{item.state}</div>
            <div class="cell">
              <button on:click={() => openUpdateModal(item)}>Update</button>
              <button on:click={() => openDeleteModal(item)}>Delete</button>
            </div>
          </div>
        {/each}
      </div>
    {:else}
      <p>Loading data...</p>
    {/if}
  
    {#if showCreationModal}
      <div class="modal">
        <div class="modal-content">
          <h3>Register CPF/CNPJ</h3>
          <input type="text" placeholder="Name" bind:value={newCpfCnpj.name} />
          <input
            type="text"
            placeholder="CPF/CNPJ"
            bind:value={newCpfCnpj.cpfNumber}
          />
          <input type="text" placeholder="City" bind:value={newCpfCnpj.city} />
          <input type="text" placeholder="State" bind:value={newCpfCnpj.state} />
          <button on:click={handleCreateCpf}>Create</button>
          <button on:click={() => (showCreationModal = false)}>Close</button>
        </div>
      </div>
    {/if}
  
    {#if showValidationModal}
      <div class="modal">
        <div class="modal-content">
          <h3>Validate CPF/CNPJ</h3>
          <input
            type="text"
            placeholder="Enter CPF"
            bind:value={validateCpfCnjpj}
          />
          <button on:click={validateInput}>Validate</button>
          {#if validationMessage}
            <div class="validation-message {validationColor}">
              {validationMessage}
            </div>
          {/if}
          <button on:click={closeValidationModal}>Close</button>
        </div>
      </div>
    {/if}
  
    {#if showUpdateModal}
      <div class="modal">
        <div class="modal-content">
          <h3>Update Item</h3>
          <input type="text" placeholder="Name" bind:value={selectedItem.name} />
          <input
            type="text"
            placeholder="CPF/CNPJ"
            bind:value={selectedItem.cpfNumber}
          />
          <input
            type="text"
            placeholder="State"
            bind:value={selectedItem.state}
          />
          <input type="text" placeholder="City" bind:value={selectedItem.city} />
          <button on:click={handleUpdateCpf}>Update</button>
          <button on:click={closeUpdateModal}>Close</button>
        </div>
      </div>
    {/if}
    {#if showDeleteModal}
      <div class="modal">
        <div class="modal-content">
          <h3>Delete Item</h3>
          <input
            type="text"
            placeholder="CPF/CNPJ"
            bind:value={selectedItem.cpfNumber}
          />
          <button on:click={handleDeleteCpf}>Delete</button>
          <button on:click={closeDeleteModal}>Close</button>
        </div>
      </div>
    {/if}
  </div>
  
  <style>
    /* Modal styles */
    .modal {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, 0.5);
      display: flex;
      justify-content: center;
      align-items: center;
    }
    .modal-content {
      background: white;
      padding: 20px;
      border-radius: 8px;
      width: 400px;
    }
    .container {
      max-width: 1200px;
      margin: 0 auto;
      padding: 16px;
      font-family: Arial, sans-serif;
    }
  
    .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
    }
  
    .actions {
      display: flex;
      gap: 10px;
      margin-bottom: 20px;
    }
  
    .filters {
      display: flex;
      gap: 10px;
      margin-bottom: 20px;
    }
  
    .table {
      display: flex;
      flex-direction: column;
      gap: 10px;
    }
  
    .row {
      display: flex;
      justify-content: space-between;
      padding: 10px;
      background-color: #f5f5f5;
      border-radius: 8px;
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
  
    .header-row {
      font-weight: bold;
      background-color: #e0e0e0;
    }
  
    .cell {
      flex: 1;
      text-align: center;
    }
  
    .blocklist {
      margin-top: 20px;
      background-color: #ffe5e5;
      padding: 10px;
      border: 1px solid #ff0000;
      border-radius: 4px;
    }
    .validation-message {
      margin-top: 8px;
      font-size: 14px;
      font-weight: bold;
    }
    .validation-message.green {
      color: green;
    }
    .validation-message.red {
      color: red;
    }
  </style>
  