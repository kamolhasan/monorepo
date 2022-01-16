# monorepo

## Related Blogs

- [golang-multimodule-monorepo-tutorial](https://irilivibi.medium.com/golang-multimodule-monorepo-tutorial-3f5cf10e9b9a)
- [catching-up-with-the-world-go-modules-in-a-monorepo](https://medium.com/compass-true-north/catching-up-with-the-world-go-modules-in-a-monorepo-c3d1393d6024)
- [go-module-a-guide-for-monorepos-part-1](https://engineering.grab.com/go-module-a-guide-for-monorepos-part-1)
- [go-module-a-guide-for-monorepos-part-2](https://engineering.grab.com/go-module-a-guide-for-monorepos-part-2)
- [full-stack-monorepo-part-i-go-services](https://medium.com/burak-tasci/full-stack-monorepo-part-i-go-services-967bb3527bb8)

## GitHub Repos

- [ianhecker/bakery](https://github.com/ianhecker/bakery)
- [flowerinthenight/golang-monorepo](https://github.com/flowerinthenight/golang-monorepo)
- [yuchanns/gobyexample](https://github.com/yuchanns/gobyexample)

## Multi-Module

- maintain dedicated go module file for each of the service 
- use `replace` in `go.mod` file to reflect the local changes in other services where it depends on.
    ```
    require (
        github.com/kamolhasan/monorepo/multi-module/service-1 v0.0.0
        github.com/kamolhasan/monorepo/multi-module/service-2 v0.0.0
    )
    
    replace (
        github.com/kamolhasan/monorepo/multi-module/service-1 => ../service-1
        github.com/kamolhasan/monorepo/multi-module/service-2 => ../service-2
    )
    ```

## Single-Module

- one root Go module file for all the services
- Less control over service level dependency version