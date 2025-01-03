<script>
    export let show = false;
    export let selectedItem = { cpfNumber: "" };
    export let handleAddCpfToBlockList;
    export let closeModal;
</script>

{#if show}
  <div class="modal">
    <div class="modal-content">
      <h3>Add Block List Item</h3>
      <input
        type="text"
        placeholder="CPF"
        bind:value={selectedItem.cpfNumber}
      />
      <button class="add-btn" on:click={handleAddCpfToBlockList}>Add</button>
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

  .add-btn {
    background-color: #4caf50;
    color: white;
  }

  .add-btn:hover {
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
