// registerParcel регистрирует новую послыку со статусом "registered"
function registerParcel() {
  const client = document.getElementById("register-client").value;
  const address = document.getElementById("register-address").value;

  if (!client || !address) {
    alert("Please provide client number and address.");
    return;
  }

  fetch(`/parcels`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ client: Number(client), address }),
  })
    .then((response) => {
      if (!response.ok) {
        return response.text().then((text) => {
          throw new Error(text);
        });
      }
      return response.text();
    })
    .then((data) => {
      console.log("Parcel Registered");
      alert("Parcel Registered successfully!");
      getParcels();
    })
    .catch((error) => {
      console.error("Error registering parcel:", error);
      alert(error.message);
    });
}

// getParcels получает список всех послок в табличном виде
function getParcels() {
  fetch(`/parcels`)
    .then((response) => response.json())
    .then((data) => {
      const parcelTable = document.getElementById("parcel-table-body");
      
      if (!Array.isArray(data) || data.length === 0) {
        console.log("No parcels found.");
        return; // Если данных нет, просто выходим из функции
      }
      
      parcelTable.innerHTML = data
        .map(
          (parcel) =>
            `<tr>
                  <td>${parcel.client}</td>
                  <td>${parcel.number}</td>
                  <td>${parcel.address}</td>
                  <td>${parcel.status}</td>
                  <td>${parcel.created_at}</td>
              </tr>`
        )
        .join("");
    })
    .catch((error) => {
      console.error("Ошибка при получении посылок:", error);
      alert(error.message);
    });
}

// toggleTable рализует скроллинг для таблицы посылок
function toggleTable() {
  const tableContainer = document.getElementById("parcel-table-container");
  const button = document.getElementById("show-parcels-button"); // Изменено на выбор правильной кнопки

  if (tableContainer.style.display === "none") {
    tableContainer.style.display = "block";
    button.textContent = "Hide Parcels"; // Изменяем текст кнопки здесь
  } else {
    tableContainer.style.display = "none";
    button.textContent = "Show Parcels"; // Изменяем текст кнопки здесь
  }

  getParcels();
}

// nextStatus изменяет статус на следующий в соответствии с предыдущим
function nextStatus() {
  const number = document.getElementById("next-status-parcel-number").value;

  if (!number) {
    alert("Please provide a parcel number.");
    return;
  }

  fetch(`/parcels/${number}/next-status`, {
    method: "PUT",
  })
    .then((response) => {
      if (!response.ok) {
        return response.text().then((text) => {
          throw new Error(text);
        });
      }
      return response.text();
    })
    .then((data) => {
      alert(`Status updated successfully!`);
      getParcels();
    })
    .catch((error) => {
      console.error("Error updating status:", error);
      alert(error.message);
    });
}

// updateAddress изменяет адрес посылки
function updateAddress() {
  const parcelNumber = document.getElementById(
    "update-address-parcel-number"
  ).value;
  const address = document.getElementById("update-address").value;

  if (!parcelNumber || !address) {
    alert("Please provide parcel number and new address.");
    return;
  }

  fetch(`/parcels/${parcelNumber}/change-address`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ address }),
  })
    .then((response) => {
      if (!response.ok) {
        return response.text().then((text) => {
          throw new Error(text);
        });
      }
      return response.text();
    })
    .then((data) => {
      console.log("Address Updated:", address);
      alert("Address Updated successfully!");
      getParcels();
    })
    .catch((error) => {
      console.error("Error updating address:", error);
      alert(error.message);
    });
}

// deleteParcel удаляет посылку
function deleteParcel() {
  const number = document.getElementById("delete-parcel-number").value;

  if (!number) {
    alert("Please provide parcel number.");
    return;
  }

  fetch(`/parcels/${number}`, {
    method: "DELETE",
  })
    .then((response) => {
      if (!response.ok) {
        return response.text().then((text) => {
          throw new Error(text);
        });
      }
      return response.text();
    })
    .then((data) => {
      console.log("Parcel Deleted:", number);
      alert("Parcel Deleted successfully!");
      getParcels();
    })
    .catch((error) => {
      console.error("Error deleting parcel:", error);
      alert(error.message);
    });
}
