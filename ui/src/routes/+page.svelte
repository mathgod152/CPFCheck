<script lang="ts">
  import { onMount } from "svelte";
  import { validateCpforCnpj } from "$lib/functions/validation";
  import type { IValidate } from "$lib/types/validateResponse";
  import {
    addCpfToBlockList,
    createeCpf,
    deleteCpf,
    InfosCpf,
    updateCpf,
  } from "$lib/functions/cpf";
  import type { ICpf } from "$lib/types/cpf";
  import BlackListModal from "../lib/components/BlocklistCpfModal.svelte";
  import {
    addCnpjToBlockList,
    createeCnpj,
    deleteCnpj,
    InfosCnpj,
    updateCnpj,
  } from "$lib/functions/cnpj";
  import type { ICnpj } from "$lib/types/cnpj";
  import CpfRow from "$lib/components/CpfRow.svelte";
  import CnpjRow from "$lib/components/CnpjRow.svelte";
  import UpdateCpfModal from "$lib/components/UpdateCpfModal.svelte";
  import UpdateCnpjModal from "$lib/components/UpdateCnpjModal.svelte";
  import DeleteCpfModal from "$lib/components/DeleteCpfModal.svelte";
  import DeleteCnpjModal from "$lib/components/DeleteCnpjModal.svelte";
  import AddCfpToBlockListModal from "$lib/components/AddCfpToBlockListModal.svelte";
  import AddCnpjToBlockListModal from "$lib/components/AddCnpjToBlockListModal.svelte";
  import BlocklistCnpjModal from "$lib/components/BlocklistCnpjModal.svelte";
  import CreateCpfModal from "$lib/components/CreateCpfModal.svelte";
  import CreateCnpjModal from "$lib/components/CreateCnpjModal.svelte";
  import ValidateCpfModal from "$lib/components/ValidateCpfModal.svelte";
  import ValidateCnpjModal from "$lib/components/ValidateCnpjModal.svelte";

  // Variáveis de estado
  let validationMessage = "";
  let validationColor = "";
  let showValidationModal = false;
  let showCreationModal = false;
  let showUpdateModal = false;
  let showDeleteModal = false;
  let newCpf: ICpf = {
    name: "",
    city: "",
    state: "",
    type: "CPF",
    cpfNumber: "",
  };
  let newCnpj: ICnpj = {
    name: "",
    city: "",
    state: "",
    type: "CNPJ",
    cnpjNumber: "",
  };
  let validateCpf = "";
  let validateCnpj = "";
  let currentType = "CPF";
  let cpfData: ICpf[] = [];
  let cnpjData: ICnpj[] = [];
  let filterState = "";
  let filterCity = "";
  let sortBy = "";
  let selectedItem: ICpf = {
    name: "",
    city: "",
    state: "",
    type: "CPF",
    cpfNumber: "",
  };
  let selectedCnpjItem: ICnpj = {
    name: "",
    city: "",
    state: "",
    type: "CNPJ",
    cnpjNumber: "",
  };
  let showAddBlockListModal = false;
  let showBlockListModal = false;
  $: filteredCpfData();
  $: filteredCnpjData();

  onMount(async () => {
    // Carrega os dados de CPF e CNPJ
    const loadedCpfData = await InfosCpf();
    cpfData = [...loadedCpfData];
    console.log("Dados de CPF carregados:", cpfData);

    const loadedCnpjData = await InfosCnpj();
    cnpjData = [...loadedCnpjData];
    console.log("Dados de CNPJ carregados:", cnpjData);
  });

  async function handleCreateCpf() {
    const isCreated = await createeCpf(newCpf);
    if (isCreated) {
      console.log("Cadastro realizado com sucesso!");
      showCreationModal = false;
      cpfData = await InfosCpf();
    } else {
      console.error("Falha ao criar cadastro.");
    }
  }
  async function handleCreateCnpj() {
    const isCreated = await createeCnpj(newCnpj);
    if (isCreated) {
      console.log("Cadastro realizado com sucesso!");
      showCreationModal = false;
      cnpjData = await InfosCnpj();
    } else {
      console.error("Falha ao criar cadastro.");
    }
  }

  async function handleUpdateCpf() {
    if (!selectedItem || !selectedItem.cpfNumber) {
      console.error("Nenhum item selecionado para atualizar.");
      return;
    }

    const isUpdated = await updateCpf(selectedItem.cpfNumber, selectedItem);
    if (isUpdated) {
      console.log("Atualização realizada com sucesso!");
      showUpdateModal = false;
    } else {
      console.error("Falha ao atualizar o cadastro.");
    }
  }
  async function handleUpdateCnpj() {
    if (!selectedCnpjItem || !selectedCnpjItem.cnpjNumber) {
      console.error("Nenhum CNPJ selecionado para atualizar.");
      return;
    }

    const isUpdated = await updateCnpj(
      selectedCnpjItem.cnpjNumber,
      selectedCnpjItem
    );
    if (isUpdated) {
      console.log("Atualização de CNPJ realizada com sucesso!");
      showUpdateModal = false;
    } else {
      console.error("Falha ao atualizar o cadastro do CNPJ.");
    }
  }

  async function handleDeleteCpf() {
    if (!selectedItem || !selectedItem.cpfNumber) {
      console.error("Nenhum item selecionado para excluir.");
      return;
    }

    const isDeleted = await deleteCpf(selectedItem.cpfNumber);
    if (isDeleted) {
      console.log("Exclusão realizada com sucesso!");
      showUpdateModal = false;
    } else {
      console.error("Falha ao excluir o cadastro.");
    }
  }
  async function handleDeleteCnpj() {
    if (!selectedCnpjItem || !selectedCnpjItem.cnpjNumber) {
      console.error("Nenhum item selecionado para excluir.");
      return;
    }

    const isDeleted = await deleteCnpj(selectedCnpjItem.cnpjNumber);
    if (isDeleted) {
      console.log("Exclusão realizada com sucesso!");
      showUpdateModal = false;
    } else {
      console.error("Falha ao excluir o cadastro.");
    }
  }

  async function handleAddCpfToBlockList() {
    if (!selectedItem || !selectedItem.cpfNumber) {
      console.error(
        "Nenhum item selecionado para adicionar à lista de bloqueio."
      );
      return;
    }

    const isAdded = await addCpfToBlockList(selectedItem.cpfNumber);
    if (isAdded) {
      console.log("CPF adicionado à lista de bloqueio com sucesso!");
      showAddBlockListModal = false;
    } else {
      console.error("Falha ao adicionar CPF à lista de bloqueio.");
    }
  }
  async function handleAddCnpjToBlockList() {
    if (!selectedCnpjItem || !selectedCnpjItem.cnpjNumber) {
      console.error(
        "Nenhum item selecionado para adicionar à lista de bloqueio."
      );
      return;
    }

    const isAdded = await addCnpjToBlockList(selectedCnpjItem.cnpjNumber);
    if (isAdded) {
      console.log("CPF adicionado à lista de bloqueio com sucesso!");
      showAddBlockListModal = false;
    } else {
      console.error("Falha ao adicionar CPF à lista de bloqueio.");
    }
  }

  function filteredCpfData() {
    let filtered = [...cpfData];
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
    if (sortBy && sortBy !== "") {
      filtered.sort((a, b) => (a[sortBy] > b[sortBy] ? 1 : -1));
    }
    return filtered;
  }

  function filteredCnpjData() {
    let filtered = [...cnpjData];
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

  function openUpdateCnpjModal(item: ICnpj) {
    selectedCnpjItem = { ...item };
    showUpdateModal = true;
  }

  function openAddBlockListModal(item: ICpf) {
    selectedItem = { ...item };
    showAddBlockListModal = true;
  }

  function openAddCnpjBlockListModal(item: ICpf) {
    selectedItem = { ...item };
    showAddBlockListModal = true;
  }

  function openDeleteModal(item: ICpf) {
    selectedItem = { ...item };
    showDeleteModal = true;
  }

  function openDeleteCnpjModal(item: ICnpj) {
    selectedCnpjItem = { ...item };
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

  function closeAddBlockListModal() {
    selectedItem = null;
    showAddBlockListModal = false;
  }

  function closeBlockListModal() {
    showBlockListModal = false;
  }

  async function validateInput() {
    const validationResponse: IValidate = await validateCpforCnpj(
      validateCpf,
      currentType
    );
    validationMessage = validationResponse.validationMessage;
    validationColor = validationResponse.validateColor;
  }
  async function validateInputCnpj() {
    const validationResponse: IValidate = await validateCpforCnpj(
      validateCnpj,
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
    validateCpf = "";
  }

  function openBlockListModal() {
    showBlockListModal = true;
  }
</script>

<div class="container">
  <div class="header">
    <h1>Validator - {currentType}</h1>
    <div class="actions">
      <button on:click={openCreationModal}>Register</button>
      <button on:click={openValidationModal}>Validate</button>
    </div>
  </div>

  <div class="filters">
    <select bind:value={sortBy}>
      <option value="">Sort By</option>
      <option value="city">City</option>
      <option value="state">State</option>
    </select>
    <button
      on:click={() => (currentType = currentType === "CPF" ? "CNPJ" : "CPF")}
    >
      Switch to {currentType === "CPF" ? "CNPJ" : "CPF"}
    </button>
    <button on:click={openBlockListModal}>
      Show {currentType} BlockList
    </button>
  </div>

  {#if (currentType == "CPF" && cpfData.length > 0) || (currentType == "CNPJ" && cnpjData.length > 0)}
    <div class="table">
      {#if currentType === "CPF"}
        <CpfRow
          {cpfData}
          {openUpdateModal}
          {openDeleteModal}
          {openAddBlockListModal}
        ></CpfRow>
      {/if}
      {#if currentType === "CNPJ"}
        <CnpjRow
          {cnpjData}
          {openUpdateCnpjModal}
          {openDeleteCnpjModal}
          {openAddCnpjBlockListModal}
        ></CnpjRow>
      {/if}
    </div>
  {:else}
    <p>Loading data...</p>
  {/if}

  {#if showCreationModal}
    {#if currentType === "CPF"}
      <CreateCpfModal
        show={showCreationModal}
        {newCpf}
        {handleCreateCpf}
        closeModal={closeCreationModal}
      />
    {/if}
    {#if currentType === "CNPJ"}
      <CreateCnpjModal
        show={showCreationModal}
        {newCnpj}
        {handleCreateCnpj}
        closeModal={closeCreationModal}
      />
    {/if}
  {/if}

  {#if showValidationModal}
    {#if currentType === "CPF"}
      <ValidateCpfModal
        bind:showValidationModal
        bind:validateCpf
        bind:validationMessage
        bind:validationColor
        {validateInput}
        {closeValidationModal}
      />
    {/if}
    {#if currentType === "CNPJ"}
      <ValidateCnpjModal
      bind:showValidationModal
      bind:validateCnpj
      bind:validationMessage
      bind:validationColor
      {validateInputCnpj}
      {closeValidationModal}
      />
    {/if}
  {/if}

  {#if showUpdateModal}
    {#if currentType === "CPF"}
      <UpdateCpfModal
        {selectedItem}
        closeModal={closeUpdateModal}
        {handleUpdateCpf}
      />
    {/if}
    {#if currentType === "CNPJ"}
      <UpdateCnpjModal
        {selectedCnpjItem}
        closeModal={closeUpdateModal}
        {handleUpdateCnpj}
      />
    {/if}
  {/if}

  {#if showDeleteModal}
    {#if currentType === "CPF"}
      <DeleteCpfModal
        title="Delete Item"
        message="Are you sure you want to delete this item?"
        onConfirm={handleDeleteCpf}
        onCancel={closeDeleteModal}
      />
    {/if}
    {#if currentType === "CNPJ"}
      <DeleteCnpjModal
        title="Delete Item"
        message="Are you sure you want to delete this item?"
        onConfirm={handleDeleteCnpj}
        onCancel={closeDeleteModal}
      />
    {/if}
  {/if}

  {#if showAddBlockListModal}
    {#if currentType === "CPF"}
      <AddCfpToBlockListModal
        show={showAddBlockListModal}
        {selectedItem}
        {handleAddCpfToBlockList}
        closeModal={closeAddBlockListModal}
      />
    {/if}
    {#if currentType === "CNPJ"}
      <AddCnpjToBlockListModal
        show={showAddBlockListModal}
        {selectedCnpjItem}
        {handleAddCnpjToBlockList}
        closeModal={closeAddBlockListModal}
      />
    {/if}
  {/if}

  {#if showBlockListModal}
    <div class="modal-blocklist">
      {#if currentType === "CPF"}
        <BlackListModal />
      {/if}
      {#if currentType === "CNPJ"}
        <BlocklistCnpjModal />
      {/if}
      <div class="modal-blocklist-butom">
        <button on:click={closeBlockListModal}>Close</button>
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
    background: rgba(192, 192, 192, 0.5);
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
  .modal-blocklist {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .modal-blocklist-butom {
    margin-top: 10px;
    width: 50%;
    height: 50px;
    justify-content: center;
    display: flex;
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
  button {
    border-radius: 5px;
    background-color: #86680c;
    color: #f5f5f5;
    border: none;
    height: 35px;
  }
</style>
