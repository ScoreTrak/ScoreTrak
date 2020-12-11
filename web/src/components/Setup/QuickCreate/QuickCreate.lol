import React, {useRef} from "react";
import Paper from "@material-ui/core/Paper";
import Box from '@material-ui/core/Box';
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepButton from "@material-ui/core/StepButton";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import makeStyles from "@material-ui/core/styles/makeStyles";
import TeamCreate from "../Team/Team";
import HostCreate from "../QuickCreate/Host";
import ServiceCreate from "../Service/Services";
import PropertiesCreate from "../Property/Properties";

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
    },
    button: {

    },
    submit: {
        display: 'inline-block',
    },
    instructions: {
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(1),
    },
    buttonsDiv: {

    },
}));

function getSteps() {
    return ['Teams', 'Hosts', 'Services', 'Properties'];
}

function getStepContent(step, props, submitRef) {
    switch (step) {
        case 0:
            return <TeamCreate {...props} />;
        case 1:
            return <HostCreate {...props} />;
        case 2:
            return <ServiceCreate {...props} />;
        case 3:
            return <PropertiesCreate {...props} />;
        default:
            return 'Unknown step';
    }
}

function getTitleContent(step: number) {
    switch (step) {
        case 0:
            return 'Step 1: Create Teams';
        case 1:
            return 'Step 2: Create Hosts';
        case 2:
            return 'Step 3: Create Services';
        case 3:
            return 'Step 4: Create Properties';
        default:
            return 'Unknown step';
    }
}

export default function QuickCreate(props) {

    const setTitle = props.setTitle
    const classesPaper = props.classesPaper
    setTitle("Quick Create")

    const classes = useStyles();
    const [activeStep, setActiveStep] = React.useState<number>(0);
    const steps = getSteps();


    const handleStep = (step: number) => () => {
        setActiveStep(step);
    };


    return (
        <Paper className={classesPaper} style={{minHeight: "85vh"}}>
            <Box height="100%" width="100%" >
                <Stepper nonLinear activeStep={activeStep}>
                    {steps.map((label, index) => (
                        <Step key={label}>
                            <StepButton onClick={handleStep(index)}>
                                {label}
                            </StepButton>
                        </Step>
                    ))}
                </Stepper>
                <div>
                    <div>
                        <Typography className={classes.instructions}>{getTitleContent(activeStep)}</Typography>
                    <div>
                        {getStepContent(activeStep, props)}
                        </div>
                    </div>
                </div>
            </Box>
        </Paper>
    );
}