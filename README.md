"## Plataforma para registrar un Dispositivo IoT cargar, escuchar u obtener las lecturas de estos dispositivos IoT" 

Estructura de carpetas


Example to register a new device:
curl --location '127.0.0.1:5000/v1/user/registerDevice' \
--header 'Content-Type: application/json' \
--data '{"ID":"ECA_0001","Name":"EstacionCalidadDelAire","Location":"BCN","Status":"True","CreatedAt":"2024-06-02T13:19:41Z","UpdatedAt":"2024-06-02T13:19:41Z"}'


