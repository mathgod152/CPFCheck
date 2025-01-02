<script>
    let data = [
      { id: 1, name: "Alice", type: "CPF", cpfCnpj: "123.456.789-00", status: "Active" },
      { id: 2, name: "Bob", type: "CNPJ", cpfCnpj: "12.345.678/0001-99", status: "Inactive" },
      { id: 3, name: "Charlie", type: "CNPJ", cpfCnpj: "98.765.432/0001-00", status: "Pending" },
      { id: 4, name: "Diana", type: "CPF", cpfCnpj: "987.654.321-00", status: "Active" },
    ];
  
    let currentType = "CPF";
    let newCpfCnpj = "";
    let filterStatus = "";
    let blocklist: any[] = [];
    let sortBy = "";
    let showUpdateModal = false;
    let selectedItem = null;
  
    function openUpdateModal(item) {
      selectedItem = { ...item };
      showUpdateModal = true;
    }
  
    function closeUpdateModal() {
      selectedItem = null;
      showUpdateModal = false;
    }
  
    function saveCpfCnpj() {
      console.log("CPF/CNPJ salvo:", newCpfCnpj);
      newCpfCnpj = "";
    }
  
    function toggleBlocklist(item) {
      if (blocklist.includes(item.id)) {
        blocklist = blocklist.filter((id) => id !== item.id);
      } else {
        blocklist = [...blocklist, item.id];
      }
    }
  
    function filteredData() {
      let filtered = data.filter((item) => item.type === currentType);
      if (filterStatus) {
        filtered = filtered.filter((item) => item.status === filterStatus);
      }
      if (sortBy) {
        filtered.sort((a, b) => (a[sortBy] > b[sortBy] ? 1 : -1));
      }
      return filtered;
    }
  </script>

    
<div class="container">
    <div class="header">
      <h1>Validator - {currentType}</h1>
    </div>
  
    <!-- Campos de Criação e Validação -->
    <div class="actions">
      <input
        type="text"
        placeholder={currentType === "CPF" ? "Enter CPF" : "Enter CNPJ"}
        bind:value={newCpfCnpj}
      />
      <button on:click={saveCpfCnpj}>Save</button>
      <button on:click={() => console.log("Validation modal opened")}>Validate</button>
    </div>
  
    <!-- Filtros e Ordenação -->
    <div class="filters">
      <select bind:value={filterStatus}>
        <option value="">All Status</option>
        <option value="Active">Active</option>
        <option value="Inactive">Inactive</option>
        <option value="Pending">Pending</option>
      </select>
      <select bind:value={sortBy}>
        <option value="">Sort By</option>
        <option value="id">ID</option>
        <option value="name">Name</option>
        <option value="status">Status</option>
      </select>
      <button on:click={() => (currentType = currentType === "CPF" ? "CNPJ" : "CPF")}>
        Switch to {currentType === "CPF" ? "CNPJ" : "CPF"}
      </button>
    </div>
  
    <!-- Tabela -->
    <div class="table">
      <div class="row header-row">
        <div class="cell">ID</div>
        <div class="cell">Name</div>
        <div class="cell">CPF/CNPJ</div>
        <div class="cell">Status</div>
        <div class="cell">Actions</div>
      </div>
  
      {#each filteredData() as item}
        <div class="row">
          <div class="cell">{item.id}</div>
          <div class="cell">{item.name}</div>
          <div class="cell">{item.cpfCnpj}</div>
          <div class="cell">{item.status}</div>
          <div class="cell">
            <input
              type="checkbox"
              checked={blocklist.includes(item.id)}
              on:change={() => toggleBlocklist(item)}
            />
            <button on:click={() => openUpdateModal(item)}>Update</button>
          </div>
        </div>
      {/each}
    </div>
  
    <!-- Blocklist -->
    {#if blocklist.length > 0}
      <div class="blocklist">
        <h3>Blocklist:</h3>
        <ul>
          {#each blocklist as id}
            <li>{data.find((item) => item.id === id).name}</li>
          {/each}
        </ul>
      </div>
    {/if}
  
    <!-- Modal de Update -->
    {#if showUpdateModal}
      <div class="modal">
        <div class="modal-content">
          <h3>Update Item</h3>
          {#if selectedItem}
            <p>
              <strong>ID:</strong> {selectedItem.id}
            </p>
            <p>
              <label>
                <strong>Name:</strong>
                <input type="text" bind:value={selectedItem.name} />
              </label>
            </p>
            <p>
              <label>
                <strong>CPF/CNPJ:</strong>
                <input type="text" bind:value={selectedItem.cpfCnpj} />
              </label>
            </p>
            <p>
              <label>
                <strong>Status:</strong>
                <select bind:value={selectedItem.status}>
                  <option value="Active">Active</option>
                  <option value="Inactive">Inactive</option>
                  <option value="Pending">Pending</option>
                </select>
              </label>
            </p>
          {/if}
          <button on:click={closeUpdateModal}>Close</button>
        </div>
      </div>
    {/if}
  </div>
  
  <style>
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
  </style>

  