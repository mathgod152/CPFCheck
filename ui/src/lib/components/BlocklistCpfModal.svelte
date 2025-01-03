<script lang="ts">
  import { getBlockListCpfs, reemoveCpfToBlockList } from "$lib/functions/cpf";
  import type { ICpf } from "$lib/types/cpf";
  import { onMount } from "svelte";

  let data: ICpf[] = [];
  onMount(async () => {
    const loadedData = await getBlockListCpfs(); 
    data = [...loadedData]; 
    console.log("Dados carregados:", data);
  });
  async function handleRemoveCpfToBlockList(item: ICpf) {
  if (!item || !item.cpfNumber) {
    console.error("Nenhum item selecionado para deletar.");
    return;
  }

  const isRemoveToBlockList = await reemoveCpfToBlockList(item.cpfNumber);
  if (isRemoveToBlockList) {
    console.log("CPF atualizado com sucesso!");
    data = data.filter((cpf) => cpf.cpfNumber !== item.cpfNumber);
  } else {
    console.error("Falha ao atualizar o CPF.");
  }
}
</script>

<div class="modal-blocklist-content">
  <h3>Block List</h3>
  <div class="content">
    {#if data.length > 0}
      <div class="table">
        <div class="row header-row">
          <div class="cell">Name</div>
          <div class="cell">CPF/CNPJ</div>
          <div class="cell">City</div>
          <div class="cell">State</div>
          <div class="cell">Actions</div>
        </div>

        {#each data as item}
          <div class="row">
            <div class="cell">{item.name}</div>
            <div class="cell">{item.cpfNumber}</div>
            <div class="cell">{item.city}</div>
            <div class="cell">{item.state}</div>
            <div class="cell">
              <button on:click={() => handleRemoveCpfToBlockList(item)}>RemoveBlockList</button>
            </div>
          </div>
        {/each}
      </div>
    {:else}
      <p>Loading data...</p>
    {/if}
  </div>
  <!--content-->
</div>

<!--modal-content-->

<style>
  .modal-blocklist-content {
    background: white;
    padding: 20px;
    border-radius: 8px;
    width: 80%;
    height: 80%;
    display: flex;
    justify-content: center;
    flex-direction: column;
  }
  .modal-blocklist-content h3 {
    display: flex;
    justify-content: center;
  }
  .content {
    display: flex;
    height: 90%;
    width: 100%;
  }
  .table {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 100%;
  }
  .row {
    display: flex;
    justify-content: space-between;
    width: 100%;
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
  button {
    border-radius: 5px;
    background-color: #f0dda4;
    border: none;
  }
</style>
