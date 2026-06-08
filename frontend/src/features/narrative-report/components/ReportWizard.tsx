/**
 * Report Wizard Component
 * Wizard for creating narrative reports step by step
 */

import { Box, Stepper, Step, StepLabel, Button, Typography, Stack } from '@mui/material';
import { useState } from 'react';

interface ReportWizardProps {
  steps: string[];
  onComplete: () => void;
  onCancel: () => void;
}

export const ReportWizard = ({ steps, onComplete, onCancel }: ReportWizardProps) => {
  const [activeStep, setActiveStep] = useState(0);

  const handleNext = () => {
    if (activeStep < steps.length - 1) {
      setActiveStep(activeStep + 1);
    } else {
      onComplete();
    }
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };

  return (
    <Box>
      <Stepper activeStep={activeStep} sx={{ mb: 4 }}>
        {steps.map((label) => (
          <Step key={label}>
            <StepLabel>{label}</StepLabel>
          </Step>
        ))}
      </Stepper>

      <Box sx={{ mb: 4 }}>
        <Typography variant="h6">{steps[activeStep]}</Typography>
        <Typography variant="body2" color="textSecondary">
          Step {activeStep + 1} of {steps.length}
        </Typography>
      </Box>

      <Stack direction="row" spacing={2} justifyContent="flex-end">
        <Button
          disabled={activeStep === 0}
          onClick={handleBack}
        >
          Back
        </Button>
        <Button
          variant="contained"
          onClick={handleNext}
        >
          {activeStep === steps.length - 1 ? 'Finish' : 'Next'}
        </Button>
        <Button onClick={onCancel} color="error">
          Cancel
        </Button>
      </Stack>
    </Box>
  );
};
