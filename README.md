# Go for spa

Este es un servidor simple que permite ejecutar aplicaciones SPA hechas con Vue/React/Angular.

Las peticiones que se realicen diferentes de `/` serán "ignoradas" para que el router del spa sirva lo que requiere.

## Instalación y ejecución

```bash
$ go get github.com/alexyslozada/go-for-spa
$ cd $GOPATH/src/github.com/alexyslozada/go-for-spa
$ go bulid
$ ./go-for-spa
```

## Configuración

En el archivo config.json se configura información como:

* public_dir: Directorio a servir. Es donde se encuentra toda la app frontend.
* port: El puerto donde quiere servirse. Este puerto no puede ser 443 ya que no está para servirse con ssl aún. Y si se elige el 80 debe ejecutarse el binario como root.
* index: Es el archivo index a servir, por lo general `index.html`.

