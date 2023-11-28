let currentPage = 1;
const limit = 10;

async function fetchJson(url) {
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return await response.json();
}

async function displayItems(page = 1) {
    try {
        const items = await fetchJson(`/api/items?page=${page}&limit=${limit}`);
        const itemsList = document.getElementById('itemsList');
        itemsList.innerHTML = '';

        items.forEach((item) => {
            const itemElement = document.createElement('div');
            itemElement.classList.add('card', 'mb-2');
            itemElement.innerHTML = `
                    <div class="card-body">
                        <h5 class="card-title">Part Number: ${item.partNumber}</h5>
                        <div class="row">
                            <div class="col-md-6">
                                <p class="card-text">Description: ${item.description}</p>
                                <p class="card-text">Price: ${item.price}</p>
                                <p class="card-text">Quantity: ${item.quantity}</p>
                            </div>
                            <div class="col-md-6">
                                <p class="card-text">Purchase Order: ${item.purchaseOrder}</p>
                                <p class="card-text">Serial Number: ${item.serialNumber}</p>
                                <p class="card-text">Category: ${item.category}</p>
                            </div>
                        </div>
                    </div>`;
            itemsList.appendChild(itemElement);
        });
    } catch (error) {
        console.error('Error:', error);
    }
}

document.getElementById('addItemForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const newItem = {
        partNumber: document.getElementById('partNumber').value,
        serialNumber: document.getElementById('serialNumber').value,
        purchaseOrder: document.getElementById('purchaseOrder').value,
        description: document.getElementById('description').value,
        category: document.getElementById('category').value,
        price: parseFloat(document.getElementById('price').value) || null,
        quantity: parseInt(document.getElementById('quantity').value) || null,
        inventoryId: '',
    };

    try {
        const data = await fetchJson('/api/items', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(newItem),
        });
        console.log('Success:', data);
        displayItems();
    } catch (error) {
        console.error('Error:', error);
    }

    this.reset();
});

document.getElementById('searchItemForm').addEventListener('submit', async function (e) {
    e.preventDefault();
    const query = document.getElementById('searchQuery').value;

    try {
        const data = await fetchJson(`/api/items?search=${encodeURIComponent(query)}`);
        const searchResults = document.getElementById('searchResults');
        searchResults.innerHTML = '';

        const items = Array.isArray(data) ? data : [data];

        items.forEach((item) => {
            const itemElement = document.createElement('div');
            itemElement.classList.add('card', 'mb-2');
            itemElement.innerHTML = `
                    <div class="card-body">
                        <h5 class="card-title">Part Number: ${item.partNumber}</h5>
                        <div class="row">
                            <div class="col-md-6">
                                <p class="card-text">Description: ${item.description}</p>
                                <p class="card-text">Price: ${item.price}</p>
                                <p class="card-text">Quantity: ${item.quantity}</p>
                            </div>
                            <div class="col-md-6">
                                <p class="card-text">Purchase Order: ${item.purchaseOrder}</p>
                                <p class="card-text">Serial Number: ${item.serialNumber}</p>
                                <p class="card-text">Category: ${item.category}</p>
                            </div>
                        </div>
                    </div>`;
            searchResults.appendChild(itemElement);
        });
    } catch (error) {
        console.error('Error:', error);
    }
});

document.getElementById('prevPage').addEventListener('click', function (e) {
    e.preventDefault();
    if (currentPage > 1) {
        currentPage--;
        displayItems(currentPage);
    }
});

document.getElementById('nextPage').addEventListener('click', function (e) {
    e.preventDefault();
    currentPage++;
    displayItems(currentPage);
});

displayItems();
