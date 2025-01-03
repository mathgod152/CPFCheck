<script lang="ts">
    export let title: string = "Delete Item";
    export let message: string = "Are you sure you want to delete this item?";
    export let onConfirm: () => void;
    export let onCancel: () => void;
  </script>
  
  <div class="modal">
    <div class="modal-content">
      <h3>{title}</h3>
      <p>{message}</p>
      <div class="actions">
        <button on:click={onConfirm}>Confirm</button>
        <button on:click={onCancel}>Cancel</button>
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
    .actions {
      display: flex;
      justify-content: space-between;
      margin-top: 20px;
    }
    button {
      border-radius: 5px;
      background-color: #86680c;
      border: none;
      height: 35px;
      padding: 0 10px;
      color: white;
      cursor: pointer;
    }
  </style>
  