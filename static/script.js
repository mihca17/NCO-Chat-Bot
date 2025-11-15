let myMap;

// Города присутствия Росатома
const cities = [
    { name: "Ангарск", region: "Иркутская область", coords: [52.5618, 103.9238], organizations : [{ orgname : "ОО ТОС АГО \"12а микрорайон\"", category : "Местное сообщество и развитие территорий", description : "Повышение качества жизни жителей 12а микрорайона г.Ангарска Иркутской области", contacts : "https://vk.com/id746471055"}] },
    { name: "Волгодонск", region: "Ростовская область", coords: [47.5142, 42.2075],organizations : [{ orgname : "Благотворительный общественно полезный фонд помощи социально незащищенным слоям населения \"Платформа добрых дел\"", category : "Социальная защита (помощь людям в трудной ситуации)", description : "Благотворительный общественно полезный фонд помощи социально незащищенным слоям населения «Платформа добрых дел» Основной вид деятельности (ОКВЭД) 64.99", contacts : "нет"}]  },
    { name: "Глазов", region: "Удмуртская Республика", coords: [58.1489, 52.6603],organizations : []  },
    { name: "Десногорск", region: "Смоленская область", coords: [54.1405, 33.3144] },
    { name: "Димитровград", region: "Ульяновская область", coords: [54.2201, 49.5701] },
    { name: "Железногорск", region: "Красноярский край", coords: [56.2304, 93.4842] },
    { name: "ЗАТО Заречный", region: "Пензенская область", coords: [53.1925, 45.1844] },
    { name: "Заречный", region: "Свердловская область", coords: [56.8008, 61.3128] },
    { name: "Зеленогорск", region: "Красноярский край", coords: [55.9511, 92.5953] },
    { name: "Краснокаменск", region: "Забайкальский край", coords: [50.0942, 118.0439] },
    { name: "Курчатов", region: "Курская область", coords: [51.8061, 35.0591] },
    { name: "Лесной", region: "Свердловская область", coords: [58.6350, 59.7770] },
    { name: "Неман", region: "Калининградская область", coords: [55.1108, 22.0334] },
    { name: "Нововоронеж", region: "Воронежская область", coords: [51.3014, 39.2214] },
    { name: "Новоуральск", region: "Свердловская область", coords: [57.2439, 60.0922] },
    { name: "Обнинск", region: "Калужская область", coords: [55.0950, 36.6138] },
    { name: "Озерск", region: "Челябинская область", coords: [55.7300, 60.7100] },
    { name: "Певек", region: "Чукотский АО", coords: [69.9883, 170.3074] },
    { name: "Полярные Зори", region: "Мурманская область", coords: [67.3750, 32.4383] },
    { name: "Саров", region: "Нижегородская область", coords: [54.9586, 43.3100] },
    { name: "Северск", region: "Томская область", coords: [56.0891, 85.5625] },
    { name: "Снежинск", region: "Челябинская область", coords: [56.0419, 60.0981] },
    { name: "Советск", region: "Калининградская область", coords: [55.0650, 21.5061] },
    { name: "Сосновый Бор", region: "Ленинградская область", coords: [59.8600, 29.0800] },
    { name: "Трехгорный", region: "Челябинская область", coords: [55.8600, 60.6400] },
    { name: "Удомля", region: "Тверская область", coords: [57.8800, 34.9600] },
    { name: "Усолье-Сибирское", region: "Иркутская область", coords: [52.7600, 103.6300] },
    { name: "Электросталь", region: "Московская область", coords: [55.7800, 38.4400] },
    { name: "Энергодар", region: "Запорожская область", coords: [47.5000, 34.6500] }
];
function initMap() {
    document.querySelector('.loading').style.display = 'none';

    myMap = new ymaps.Map('map', {
        center: [55.7558, 37.6176], // Москва по умолчанию
        zoom: 5,
        controls: ['zoomControl', 'fullscreenControl']
    });

    // Добавляем все метки
    cities.forEach(city => {
        const placemark = new ymaps.Placemark(
            city.coords,
            {
                balloonContent: `<b>${city.name}</b>, ${city.region}`,
                hintContent: `${city.name}`
            },
            {
                preset: 'islands#blueCircleIcon'
            }
        );

        myMap.geoObjects.add(placemark);

        // При клике на метку открываем карточку НКО с информацией об организациях
        placemark.events.add('click', function () {
            displayOrgCard(city.organizations);
        });
    });
}

// Функция для отображения информации об организациях в карточке
function displayOrgCard(orgs) {
    const card = document.getElementById('ngoCard');
    const cardContent = document.querySelector('#ngoCard .card-content');

    if (!cardContent) {
        // Если элемента .card-content нет, создадим его
        card.innerHTML = `
            <span class="close-btn">&times;</span>
            <div class="card-content"></div>
        `;
    }

    let contentHTML = '';
    orgs.forEach(org => {
        contentHTML += `
            <div class="org-item">
                <h3>${org.orgname}</h3>
                <p><strong>Категория:</strong> ${org.category}</p>
                <p><strong>Описание:</strong> ${org.description}</p>
                <p><strong>Контакты:</strong> <a href="${org.contacts}" target="_blank">${org.contacts}</a></p>
            </div>
            <hr>
        `;
    });

    document.querySelector('#ngoCard .card-content').innerHTML = contentHTML;
    card.style.display = 'block';
}

document.querySelector('.close-btn').addEventListener('click', () => {
    document.getElementById('ngoCard').style.display = 'none';
});

document.querySelectorAll('.filter-btn').forEach(btn => {
    btn.addEventListener('click', () => {
        document.querySelectorAll('.filter-btn').forEach(b => b.classList.remove('active'));
        btn.classList.add('active');
    });
});

ymaps.ready(initMap);