# ScoreTrak Web App

The react web application for ScoreTrak, a scoring engine built in go.

## Installation

Clone the repository.

```bash
git clone https://github.com/ScoreTrak/ScoreTrak
```

## Usage

Navigate to the web application

```bash
cd ScoreTrak/web
```

Install Dependencies

```bash
npm install
```

### Either build, test, and/or start the application

Build the application

```bash
npm run build
```

Test the application

```bash
npm run test
```

Start the application

```bash
npm run start
```

## Notes

1. If you are trying to install dependcies and can't with internet access, you may need to configure a proxy

```bash
# https://docs.npmjs.com/cli/v7/using-npm/config#https-proxy
npm config set proxy http://<username>:<password>@<proxy-server-url>:<port>

# https://docs.npmjs.com/cli/v7/using-npm/config#proxy
npm config set https-proxy http://<username>:<password>@<proxy-server-url>:<port>
```

Thanks to [npm proxy guide.](https://www.ngdevelop.tech/npm-proxy-setting/)

2. This application templated from a react + typescript + grpc