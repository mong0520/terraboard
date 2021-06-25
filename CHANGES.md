# Changelog

## [1.2.0](https://www.github.com/camptocamp/terraboard/compare/v1.1.0...v1.2.0) (2021-06-25)


### Features

* add plan db scheme ([#164](https://www.github.com/camptocamp/terraboard/issues/164)) ([2c1cd8f](https://www.github.com/camptocamp/terraboard/commit/2c1cd8f003016ee3c54a57cf6d679e6ee2fc5b29))
* add test environment with proper docker-compose using MinIO + add compatibility for it to AWS provider ([#165](https://www.github.com/camptocamp/terraboard/issues/165)) ([6d44ecb](https://www.github.com/camptocamp/terraboard/commit/6d44ecbf5765e33b2cfcdbef1dce9059fe7dc94a))
* **api:** add lineages get endpoint ([#176](https://www.github.com/camptocamp/terraboard/issues/176)) ([f632586](https://www.github.com/camptocamp/terraboard/commit/f6325861cca0306382ccfd897740fa25911fc27c))
* **api:** add plan submit/get endpoints   ([#175](https://www.github.com/camptocamp/terraboard/issues/175)) ([f341ffd](https://www.github.com/camptocamp/terraboard/commit/f341ffdb0c990414e349a3596221dbdd8d4668e9))
* **gorm:** gorm v2 migration ([#170](https://www.github.com/camptocamp/terraboard/issues/170)) ([630ce3f](https://www.github.com/camptocamp/terraboard/commit/630ce3fe4044f44d66dad854430de3bee3412591))
* new standalone lineage table + associated migration ([#173](https://www.github.com/camptocamp/terraboard/issues/173)) ([aa7d455](https://www.github.com/camptocamp/terraboard/commit/aa7d455cf9ef86651ad0791b02747502b8a44c4e))


### Bug Fixes

* terraboard crash at compose startup if db isn't fully initialized ([#174](https://www.github.com/camptocamp/terraboard/issues/174)) ([a8e6bdd](https://www.github.com/camptocamp/terraboard/commit/a8e6bdd7a9195dd54af8079a4d9c1101e6cadbb9))
