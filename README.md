# Автозвонилка с API

Запуск контейнера:
```
./docker/build_image.sh
```
Конфиг: 
- config/config.xml

Как работает:
- GET http://localhost:8080/v1/call - запустит звонок с номера VP_ENV_CALLER на номера: VP_ENV_CALLEE_*
- Если подняли трубку - отправляет GET http://localhost:3540/callbot?result=true (API_FALSE_RESULT)
- Во всех остальных случаях - GET http://localhost:3540/callbot?result=false (API_FALSE_RESULT)

## Environment

- VP_ENV_CALLER= номер с которого будем звонить (например: 749910101001101@127.0.0.1)
- VP_ENV_CALLEE_*= номера на которые будем звонить (например: 7499101010101@login.mtt.ru)
- VP_ENV_ACCOUNT= номер акк. у провайдера
- VP_ENV_USERNAME= имя пользователя у провайдера (может совпадать с VP_ENV_ACCOUNT)
- VP_ENV_PASSWORD= пароль
- VP_ENV_REALM= адрес провайдера (например login.mtt.ru)
- VP_ENV_REGISTRAR= адрес провайдера (может сопадать с VP_ENV_REALM)