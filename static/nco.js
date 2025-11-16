let myMap;
let selectedPlacemark = null;
let selectedCoords = null;

// Инициализация карты
ymaps.ready(initMap);

function initMap() {
    myMap = new ymaps.Map('map', {
        center: [55.7558, 37.6176], // Москва по умолчанию
        zoom: 10,
        controls: ['zoomControl', 'fullscreenControl']
    });

    // Обработчик клика по карте
    myMap.events.add('click', function (e) {
        const coords = e.get('coords');
        setSelectedCoordinates(coords);
        createPlacemark(coords);
    });

    // Загрузка существующих организаций (если нужно)
    loadExistingOrganizations();
}

function setSelectedCoordinates(coords) {
    selectedCoords = coords;
    const coordinatesDisplay = document.getElementById('coordinatesDisplay');
    const coordX = document.getElementById('coordX');
    const coordY = document.getElementById('coordY');
    const submitBtn = document.getElementById('submitBtn');
    
    coordinatesDisplay.textContent = `Широта: ${coords[0].toFixed(6)}, Долгота: ${coords[1].toFixed(6)}`;
    coordinatesDisplay.parentElement.classList.add('has-coordinates');
    
    coordX.value = coords[0];
    coordY.value = coords[1];
    console.log(coordX.value,coordY.value)

    // Активируем кнопку отправки
    submitBtn.disabled = false;
}

function createPlacemark(coords) {
    // Удаляем предыдущую метку
    if (selectedPlacemark) {
        myMap.geoObjects.remove(selectedPlacemark);
    }
    
    // Создаем новую метку
    selectedPlacemark = new ymaps.Placemark(coords, {
        balloonContent: 'Выбранное местоположение для новой организации',
        hintContent: 'Новая организация'
    }, {
        preset: 'islands#blueIcon',
        draggable: true
    });
    
    myMap.geoObjects.add(selectedPlacemark);
    
    // Обработчик перемещения метки
    selectedPlacemark.events.add('dragend', function () {
        const newCoords = selectedPlacemark.geometry.getCoordinates();
        setSelectedCoordinates(newCoords);
    });
}

function loadExistingOrganizations() {
    // Здесь можно загрузить существующие организации с сервера
    // Пока оставляем пустым для демонстрации
}

// Обработчики событий
document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('ncoForm');
    const clearCoordsBtn = document.getElementById('clearCoordinates');
    const cancelBtn = document.getElementById('cancelBtn');
    const successModal = document.getElementById('successModal');
    const closeModal = document.querySelector('.close-modal');
    const addAnotherBtn = document.getElementById('addAnother');
    const goToMapBtn = document.getElementById('goToMap');

    // Очистка координат
    clearCoordsBtn.addEventListener('click', function() {
        if (selectedPlacemark) {
            myMap.geoObjects.remove(selectedPlacemark);
            selectedPlacemark = null;
        }
        selectedCoords = null;
        
        const coordinatesDisplay = document.getElementById('coordinatesDisplay');
        const coordX = document.getElementById('coordX');
        const coordY = document.getElementById('coordY');
        const submitBtn = document.getElementById('submitBtn');
        
        coordinatesDisplay.textContent = 'Кликните на карте для выбора местоположения';
        coordinatesDisplay.parentElement.classList.remove('has-coordinates');
        coordX.value = '';
        coordY.value = '';
        submitBtn.disabled = true;
    });

    // Отмена
    cancelBtn.addEventListener('click', function() {
        if (confirm('Вы уверены, что хотите отменить добавление? Введенные данные будут потеряны.')) {
            window.location.href = 'index.html';
        }
    });

    // Отправка формы
    form.addEventListener('submit', function(e) {
        e.preventDefault();
        
        if (!selectedCoords) {
            alert('Пожалуйста, выберите местоположение на карте');
            return;
        }
        
        const formData = {
            name: document.getElementById('orgName').value,
            x: selectedCoords[0],
            y: selectedCoords[1],
            category: document.getElementById('category').value,
            description: document.getElementById('description').value,
            contacts: document.getElementById('contacts').value,
            city: document.getElementById('city').value,
            region: document.getElementById('region').value
        };
        
        // Здесь будет отправка данных на сервер
        saveOrganization(formData);
    });

    // Модальное окно
    closeModal.addEventListener('click', function() {
        successModal.style.display = 'none';
    });

    addAnotherBtn.addEventListener('click', function() {
        successModal.style.display = 'none';
        form.reset();
        clearCoordsBtn.click();
        myMap.setBounds(myMap.getBounds()); // Сброс вида карты
    });

    goToMapBtn.addEventListener('click', function() {
        window.location.href = 'index.html';
    });

    // Закрытие модального окна при клике вне его
    window.addEventListener('click', function(e) {
        if (e.target === successModal) {
            successModal.style.display = 'none';
        }
    });
});

// Функция сохранения организации
function saveOrganization(data) {
    console.log('Данные для сохранения:', data);

    fetch(`http://localhost:8080/api/nco`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(result => {
        showSuccessModal();
    })
    .catch(error => {
        console.error('Ошибка:', error);
        alert('Произошла ошибка при сохранении');
    });
}

function showSuccessModal() {
    const modal = document.getElementById('successModal');
    modal.style.display = 'flex';
}

// Автозаполнение региона по городу (простая реализация)
document.getElementById('city').addEventListener('blur', function() {
    const city = this.value.toLowerCase();
    const regionField = document.getElementById('region');
    
    if (regionField.value) return; // Не перезаписываем, если уже заполнено
    
    const cityRegionMap = {
        'москва': 'Московская область',
        'санкт-петербург': 'Ленинградская область',
        'екатеринбург': 'Свердловская область',
        'новосибирск': 'Новосибирская область',
        'казань': 'Республика Татарстан',
        'нижний новгород': 'Нижегородская область',
        'челябинск': 'Челябинская область',
        'самара': 'Самарская область',
        'омск': 'Омская область',
        'ростов-на-дону': 'Ростовская область',
        'уфа': 'Республика Башкортостан',
        'красноярск': 'Красноярский край',
        'пермь': 'Пермский край',
        'воронеж': 'Воронежская область',
        'волгоград': 'Волгоградская область'
    };
    
    if (cityRegionMap[city]) {
        regionField.value = cityRegionMap[city];
    }
});