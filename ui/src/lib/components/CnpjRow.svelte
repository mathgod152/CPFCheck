<script lang="ts">
    import type { ICnpj } from "$lib/types/cnpj";
  
    export let cnpjData: ICnpj[] = [];
    export let openUpdateCnpjModal: (item: ICnpj) => void;
    export let openDeleteCnpjModal: (item: ICnpj) => void;
    export let openAddCnpjBlockListModal: (item: ICnpj) => void;
  </script>
  
  <div class="table">
    <div class="row header-row">
      <div class="cell">Name</div>
      <div class="cell">CNPJ</div>
      <div class="cell">City</div>
      <div class="cell">State</div>
      <div class="cell">Actions</div>
    </div>
    
    {#each cnpjData as item}
      <div class="row">
        <div class="cell">{item.name}</div>
        <div class="cell">{item.cnpjNumber}</div>
        <div class="cell">{item.city}</div>
        <div class="cell">{item.state}</div>
        <div class="cell">
          <button on:click={() => openUpdateCnpjModal(item)}>Update</button>
          <button on:click={() => openDeleteCnpjModal(item)}>Delete</button>
          <button on:click={() => openAddCnpjBlockListModal(item)}>
            Add BlockList
          </button>
        </div>
      </div>
    {/each}
  </div>
  
  <style>
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
  
    button {
      border-radius: 5px;
      background-color: #86680c;
      border: none;
      height: 35px;
      color: white;
      cursor: pointer;
    }
  
    button:hover {
      background-color: #a9820d;
    }
  </style>
  