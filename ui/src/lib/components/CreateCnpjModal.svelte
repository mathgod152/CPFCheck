<script lang="ts">
  import type { ICnpj } from "$lib/types/cnpj";

  export let show = false;
  export let newCnpj: ICnpj;
  export let handleCreateCnpj;
  export let closeModal;
</script>

{#if show}
  <div class="modal">
    <div class="modal-content">
      <h3>Register Cnpj/CNPJ</h3>
      <input type="text" placeholder="Name" bind:value={newCnpj.name} />
      <input
        type="text"
        placeholder="Cnpj/CNPJ"
        bind:value={newCnpj.cnpjNumber}
      />
      <input type="text" placeholder="City" bind:value={newCnpj.city} />
      <input type="text" placeholder="State" bind:value={newCnpj.state} />
      <button class="create-btn" on:click={handleCreateCnpj}>Create</button>
      <button class="close-btn" on:click={closeModal}>Close</button>
    </div>
  </div>
{/if}

<style>
  .modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  .modal-content {
    background-color: #fff;
    padding: 20px;
    border-radius: 8px;
    width: 90%;
    max-width: 400px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    text-align: center;
  }

  h3 {
    margin-bottom: 20px;
    font-size: 1.5rem;
  }

  input {
    width: calc(100% - 20px);
    margin: 10px 0;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 1rem;
  }

  button {
    padding: 10px 20px;
    margin: 10px 5px;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    cursor: pointer;
  }

  .create-btn {
    background-color: #4caf50;
    color: white;
  }

  .create-btn:hover {
    background-color: #45a049;
  }

  .close-btn {
    background-color: #f44336;
    color: white;
  }

  .close-btn:hover {
    background-color: #e41f1f;
  }
</style>
