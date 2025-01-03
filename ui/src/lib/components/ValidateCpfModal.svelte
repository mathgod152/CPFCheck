<script>
    export let showValidationModal = false;
    export let validateCpf = "";
    export let validationMessage = "";
    export let validationColor = "";
    export let validateInput;
    export let closeValidationModal;
  </script>
  
  {#if showValidationModal}
    <div class="modal">
      <div class="modal-content">
        <h3>Validate CPF</h3>
        <input
          type="text"
          placeholder="Enter CPF"
          bind:value={validateCpf}
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
  
    .validation-message {
      margin-top: 10px;
      padding: 10px;
      border-radius: 4px;
      font-size: 1rem;
    }
  
    .validation-message.green {
      background-color: #4caf50;
      color: white;
    }
  
    .validation-message.red {
      background-color: #f44336;
      color: white;
    }
  </style>
  