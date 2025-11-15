ymaps.ready(init);

function init() {
    console.log('✅ Яндекс.Карты успешно загружены');
    
    // Создаем карту
    var map = new ymaps.Map('map', {
        center: [56.85, 53.22], // Центр России
        zoom: 4,
        controls: ['zoomControl', 'fullscreenControl', 'typeSelector', 'rulerControl']
    });

    // Добавляем поведение карты
    map.behaviors.enable(['scrollZoom', 'dblClickZoom']);

    // Города-миллионники России с более подробными данными
    var millionCities = [
        {
            name: 'Москва',
            coords: [55.7558, 37.6173],
            population: '12.7 млн',
            founded: '1147 г.',
            area: '2561 км²',
            color: 'islands#redIcon'
        },
        {
            name: 'Санкт-Петербург',
            coords: [59.9343, 30.3351],
            population: '5.6 млн',
            founded: '1703 г.',
            area: '1439 км²',
            color: 'islands#blueIcon'
        },
        {
            name: 'Новосибирск',
            coords: [55.0084, 82.9357],
            population: '1.6 млн',
            founded: '1893 г.',
            area: '505 км²',
            color: 'islands#darkOrangeIcon'
        },
        {
            name: 'Екатеринбург',
            coords: [56.8389, 60.6057],
            population: '1.5 млн',
            founded: '1723 г.',
            area: '495 км²',
            color: 'islands#darkOrangeIcon'
        },
        {
            name: 'Казань',
            coords: [55.7961, 49.1064],
            population: '1.3 млн',
            founded: '1005 г.',
            area: '614 км²',
            color: 'islands#greenIcon'
        },
        {
            name: 'Нижний Новгород',
            coords: [56.3269, 44.0065],
            population: '1.2 млн',
            founded: '1221 г.',
            area: '466 км²',
            color: 'islands#greenIcon'
        },
        {
            name: 'Челябинск',
            coords: [55.1644, 61.4368],
            population: '1.2 млн',
            founded: '1736 г.',
            area: '530 км²',
            color: 'islands#greenIcon'
        },
        {
            name: 'Красноярск',
            coords: [56.0153, 92.8932],
            population: '1.2 млн',
            founded: '1628 г.',
            area: '379 км²',
            color: 'islands#greenIcon'
        },
        {
            name: 'Самара',
            coords: [53.1959, 50.1002],
            population: '1.1 млн',
            founded: '1586 г.',
            area: '541 км²',
            color: 'islands#violetIcon'
        },
        {
            name: 'Уфа',
            coords: [54.7355, 55.9587],
            population: '1.1 млн',
            founded: '1574 г.',
            area: '708 км²',
            color: 'islands#violetIcon'
        },
        {
            name: 'Ростов-на-Дону',
            coords: [47.2225, 39.7188],
            population: '1.1 млн',
            founded: '1749 г.',
            area: '348 км²',
            color: 'islands#violetIcon'
        },
        {
            name: 'Омск',
            coords: [54.9924, 73.3686],
            population: '1.1 млн',
            founded: '1716 г.',
            area: '567 км²',
            color: 'islands#violetIcon'
        },
        {
            name: 'Краснодар',
            coords: [45.0355, 38.9753],
            population: '1.1 млн',
            founded: '1793 г.',
            area: '841 км²',
            color: 'islands#violetIcon'
        },
        {
            name: 'Воронеж',
            coords: [51.6720, 39.1843],
            population: '1.1 млн',
            founded: '1586 г.',
            area: '596 км²',
            color: 'islands#violetIcon'
        },
        {
            name: 'Пермь',
            coords: [58.0105, 56.2502],
            population: '1.0 млн',
            founded: '1723 г.',
            area: '800 км²',
            color: 'islands#orangeIcon'
        },
        {
            name: 'Волгоград',
            coords: [48.7194, 44.5018],
            population: '1.0 млн',
            founded: '1589 г.',
            area: '859 км²',
            color: 'islands#orangeIcon'
        }
    ];

    // Создаем кластер для меток
    var clusterer = new ymaps.Clusterer({
        clusterDisableClickZoom: true,
        clusterOpenBalloonOnClick: true,
        clusterBalloonContentLayout: 'cluster#balloonTwoColumns',
        clusterBalloonPanelMaxMapArea: 0,
        clusterBalloonContentLayoutWidth: 300,
        clusterBalloonContentLayoutHeight: 200,
        clusterBalloonPagerSize: 5
    });

    // Добавляем метки для каждого города
    millionCities.forEach(function(city, index) {
        var placemark = new ymaps.Placemark(city.coords, {
            balloonContentHeader: `<strong>${city.name}</strong>`,
            balloonContentBody: `
                <div class="balloon">
                    <p><strong>Население:</strong> ${city.population}</p>
                    <p><strong>Основан:</strong> ${city.founded}</p>
                    <p><strong>Площадь:</strong> ${city.area}</p>
                    <p><strong>Регион:</strong> ${getRegion(city.name)}</p>
                </div>
            `,
            balloonContentFooter: '<em>Город-миллионник России</em>',
            hintContent: `${city.name} - ${city.population}`
        }, {
            preset: city.color,
            balloonCloseButton: true,
            hideIconOnBalloonOpen: false
        });

        clusterer.add(placemark);
    });

    // Добавляем кластер на карту
    map.geoObjects.add(clusterer);

    // Функция для определения региона
    function getRegion(cityName) {
        var regions = {
            'Москва': 'Центральный федеральный округ',
            'Санкт-Петербург': 'Северо-Западный федеральный округ',
            'Новосибирск': 'Сибирский федеральный округ',
            'Екатеринбург': 'Уральский федеральный округ',
            'Казань': 'Приволжский федеральный округ',
            'Нижний Новгород': 'Приволжский федеральный округ',
            'Челябинск': 'Уральский федеральный округ',
            'Красноярск': 'Сибирский федеральный округ',
            'Самара': 'Приволжский федеральный округ',
            'Уфа': 'Приволжский федеральный округ',
            'Ростов-на-Дону': 'Южный федеральный округ',
            'Омск': 'Сибирский федеральный округ',
            'Краснодар': 'Южный федеральный округ',
            'Воронеж': 'Центральный федеральный округ',
            'Пермь': 'Приволжский федеральный округ',
            'Волгоград': 'Южный федеральный округ'
        };
        return regions[cityName] || 'Россия';
    }

    // Подгоняем карту чтобы были видны все метки
    map.setBounds(clusterer.getBounds(), {
        checkZoomRange: true,
        zoomMargin: 50
    });

    // Обновляем статистику
    document.getElementById('stats').textContent = 
        `Всего городов-миллионников: ${millionCities.length} • Общее население: ≈30 млн человек`;

    console.log(`🗺️ Карта успешно инициализирована с ${millionCities.length} городами`);
}