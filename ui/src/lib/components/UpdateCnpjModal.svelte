<script lang="ts">
  import type { ICnpj } from "$lib/types/cnpj";


  export let selectedCnpjItem: ICnpj;
  export let closeModal: () => void;
  export let handleUpdateCnpj: () => void;
</script>

<div class="modal">
  <div class="modal-content">
    <h3>Update CNPJ</h3>
    <label>
      Name:
      <input
        type="text"
        placeholder="Name"
        bind:value={selectedCnpjItem.name}
      />
    </label>
    <label>
      CNPJ:
      <input
        type="text"
        placeholder="CNPJ"
        bind:value={selectedCnpjItem.cnpjNumber}
      />
    </label>
    <label>
      State:
      <input
        type="text"
        placeholder="State"
        bind:value={selectedCnpjItem.state}
      />
    </label>
    <label>
      City:
      <input
        type="text"
        placeholder="City"
        bind:value={selectedCnpjItem.city}
      />
    </label>
    <div class="modal-actions">
      <button class="update-btn" on:click={handleUpdateCnpj}>Update</button>
      <button class="close-btn" on:click={closeModal}>Close</button>
    </div>
  </div>
</div>

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

  .update-btn {
    background-color: #2196f3;
    color: white;
  }

  .update-btn:hover {
    background-color: #1e88e5;
  }

  .close-btn {
    background-color: #f44336;
    color: white;
  }

  .close-btn:hover {
    background-color: #e41f1f;
  }
</style>
