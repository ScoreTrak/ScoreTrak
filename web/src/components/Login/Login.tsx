import * as React from 'react';
import {useState} from 'react';
import {Avatar} from '@material-ui/core';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import {makeStyles} from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import {Alert} from '@material-ui/lab';
import CircularProgress from '@material-ui/core/CircularProgress';
import {useForm} from "react-hook-form";
import {Severity} from "../../types/types";
import {LoginRequest} from "../../grpc/pkg/auth/auth_pb";
import {useHistory} from "react-router";
import {token} from "../../grpc/token/token";
import {AuthServiceClient} from "../../grpc/pkg/auth/AuthServiceClientPb";


const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%',
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
    spinner: {
        margin: theme.spacing(1),
    }
}));

type LoginInfo = {
    username: string;
    password: string;
};

type LoginAlertType = {
    message: string
    severity: Severity | undefined
}

type LoginProps = {
    authClient: AuthServiceClient
}

const Login = (props: LoginProps) => {
    const classes = useStyles();
    const { register, handleSubmit, errors} = useForm<LoginInfo>();
    const [loading, setLoading] = useState(false);
    const [alert, setAlert] = useState<LoginAlertType>({message: "", severity: undefined})
    const history = useHistory();
    const handleLogin = (data: LoginInfo) => {
        setAlert({severity: undefined, message: ""});
        setLoading(true);
        const loginRequest = new LoginRequest()
        loginRequest.setUsername(data.username)
        loginRequest.setPassword(data.password)
        props.authClient.login(loginRequest, {}).then(r => {
            token.saveToken(r.getAccessToken())
            history.push("/");
            history.go(0)
        },
        (error) => {
            setAlert({severity: Severity.Error, message: `Failed to login: ${error.message}.(Code: ${error.code})`});
            setLoading(false);
        })
    };

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                {
                    loading ? <CircularProgress className={classes.spinner} /> :
                    <Avatar className={classes.avatar}>
                        <LockOutlinedIcon />
                    </Avatar>
                }
                <Typography component="h1" variant="h5">
                    Sign in
                </Typography>
                <form className={classes.form} onSubmit={handleSubmit(handleLogin)}>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        inputRef={register}
                        id="username"
                        label="Username"
                        name="username"
                        autoComplete="username"
                        autoFocus
                        error={!!errors.username}
                    />
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        inputRef={register}
                        name="password"
                        label="Password"
                        type="password"
                        id="password"
                        autoComplete="current-password"
                        error={!!errors.password}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        className={classes.submit}
                    >
                        Sign In
                    </Button>
                    {alert.message && alert.severity && (
                    <Alert severity={alert.severity}>{alert.message}</Alert>
                    )}
                </form>
            </div>
        </Container>
    );
}

export default Login