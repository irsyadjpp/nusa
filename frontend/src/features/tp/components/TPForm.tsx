/**
 * TP Form Component
 * Form for creating and editing Teaching Plans
 */

import { Box, TextField, Button, Stack, Typography } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { CreateTPRequest, UpdateTPRequest } from '@/api/tp';

interface TPFormProps {
  initialValues?: Partial<CreateTPRequest>;
  onSubmit: (values: CreateTPRequest | UpdateTPRequest) => void;
  onCancel?: () => void;
  isEdit?: boolean;
}

const validationSchema = Yup.object({
  title: Yup.string().required('Title is required'),
  learning_objectives: Yup.string().required('Learning objectives are required'),
  estimated_weeks: Yup.number().required('Estimated weeks is required').min(1),
});

export const TPForm = ({ initialValues, onSubmit, onCancel, isEdit = false }: TPFormProps) => {
  const defaultValues: CreateTPRequest = {
    tp_set_id: '',
    sequence_number: 1,
    cp_id: '',
    subject_id: '',
    phase_id: '',
    element_id: '',
    subelement_id: '',
    title: '',
    learning_objectives: {
      main_objective: '',
      supporting_objectives: [],
    },
    time_allocation: {
      total_hours: 2,
      per_week_hours: 2,
      hours_per_week: 2,
      breakdown: {},
    },
    prerequisites: {
      required_tps: [],
      required_skills: [],
      notes: '',
    },
    estimated_weeks: 1,
    success_criteria: {
      mastery_thresholds: [],
      performance_indicators: [],
      minimum_requirements: [],
      minimum_mastery_level: 75,
    },
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit TP' : 'Create New TP'}
      </Typography>
      <Formik
        initialValues={{ ...defaultValues, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values: CreateTPRequest, helpers: FormikHelpers<CreateTPRequest>) => {
          onSubmit(values);
          helpers.resetForm();
        }}
      >
        {({ values, handleChange, touched, errors, isSubmitting }) => (
          <Form>
            <Stack spacing={3}>
              <TextField
                fullWidth
                label="Title"
                name="title"
                value={values.title}
                onChange={handleChange}
                error={touched.title && Boolean(errors.title)}
                helperText={touched.title && errors.title}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Learning Objectives"
                name="learning_objectives"
                value={values.learning_objectives}
                onChange={handleChange}
                error={touched.learning_objectives && Boolean(errors.learning_objectives)}
                helperText={touched.learning_objectives && typeof errors.learning_objectives === 'string' ? errors.learning_objectives : undefined}
              />

              <TextField
                fullWidth
                label="Time Allocation"
                name="time_allocation"
                value={values.time_allocation}
                onChange={handleChange}
              />

              <TextField
                fullWidth
                multiline
                rows={3}
                label="Prerequisites"
                name="prerequisites"
                value={values.prerequisites}
                onChange={handleChange}
              />

              <TextField
                fullWidth
                type="number"
                label="Estimated Weeks"
                name="estimated_weeks"
                value={values.estimated_weeks}
                onChange={handleChange}
                error={touched.estimated_weeks && Boolean(errors.estimated_weeks)}
                helperText={touched.estimated_weeks && errors.estimated_weeks}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Success Criteria"
                name="success_criteria"
                value={values.success_criteria}
                onChange={handleChange}
              />

              <Stack direction="row" spacing={2} justifyContent="flex-end">
                {onCancel && (
                  <Button variant="outlined" onClick={onCancel} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Saving...' : isEdit ? 'Update TP' : 'Create TP'}
                </Button>
              </Stack>
            </Stack>
          </Form>
        )}
      </Formik>
    </Box>
  );
};
