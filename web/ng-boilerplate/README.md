# Prysmatic Angular Boilerplate

![Node.js CI](https://github.com/prysmaticlabs/ng-boilerplate/workflows/Node.js%20CI/badge.svg)

## Background

This project serves as the backbone for all Prysmatic Labs user-facing websites, built with Angular as a standard across our company for maximum reusability, tests, and maintainability. This boilerplate includes the following

- [Tailwind CSS](https://tailwindcss.com) for a utility-first, extensible CSS framework
- [Angular Material UI](https://material.angular.io/) as a design language with useful components for all our applications
- [SCSS Lang](https://sass-lang.com/) as our css preprocessor for variables, mixins, and other useful techniques
- Auth guards and JWT interceptors as easy ways to add authentication to our applications
- Environment injection via a global service for easy retrieval of global env values
- An opinionated module system that allows for maintainability across all our projects

## Module structure and best practices

Modules in the repository are structured as follows

```
src/
  modules/
    core/
      guards/
      interceptors/
      services/
    landing/
      components/
      pages/
        home/
          home.component.ts
          home.component.spec.ts
          home.component.html
    shared/
      components/
styles/
  example.scss
  ...
```

- The **CoreModule** contains global services, guards, and interceptors used by the entire application, such as an authentication service or a JWT interceptor service to add a bearer token to all outbound API requests
- The **SharedModule** contains all shared, reusable, _small_ components for usage in the application, such as loading indicators
- Regular modules contain both a **pages** and a **components** folder, which are used for page skeletons and for individual components respectively
- Components **do not** have a NAME.component.scss file, as instead we include all our application styles in a top-level `styles/` folder so we can best leverage variables, mixins, imports, and have a neat structure to all our css code

## Installing and running

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 9.0.6 and [Node.js version 14](https://nodejs.org/en/download/). Once you have the dependencies, you can install all node modules for the project with

```
npm install
```

### Development server

Run `npm start` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

### Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

### Build

Run `npm run build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

### Running unit tests

Run `npm test` to execute the unit tests via [Karma](https://karma-runner.github.io).

### Running end-to-end tests

Run `npm run e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

