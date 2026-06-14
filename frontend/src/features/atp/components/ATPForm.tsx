/**
 * ATP Form Component
 * Form for creating and editing ATPs
 */

import { Box, TextField, Button, Stack, Typography } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { ATPCreateRequest, ATPUpdateRequest } from '@/api/atp';

interface ATPFormProps {
  initialValues?: Partial<ATPCreateRequest>;
  onSubmit: (values: ATPCreateRequest | ATPUpdateRequest) => void;
  onCancel?: () => void;
  isEdit?: boolean;
}

const validationSchema = Yup.object({
  atp_set_id: Yup.string().required('ATP Set is required'),
  tp_id: Yup.string().required('TP is required'),
  sequence_number: Yup.number().required('Sequence number is required').min(1),
  week_number: Yup.number().required('Week number is required').min(1),
  estimated_hours: Yup.number().required('Estimated hours is required').min(1),
});

export const ATPForm = ({ initialValues, onSubmit, onCancel, isEdit = false }: ATPFormProps) => {
  const defaultValues: ATPCreateRequest = {
    atp_set_id: '',
    tp_id: '',
    sequence_number: 1,
    week: 1,
    learning_activities: {
      opening: [],
      core_activities: [],
      closing: [],
    },
    assessment_methods: [],
    time_allocation: {
      total_hours: 2,
      per_week_hours: 2,
      hours_per_week: 2,
      breakdown: {},
    },
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit ATP' : 'Create New ATP'}
      </Typography>
      <Formik
        initialValues={{ ...defaultValues, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values: ATPCreateRequest, helpers: FormikHelpers<ATPCreateRequest>) => {
          onSubmit(values);
          helpers.resetForm();
        }}
      >
        {({ values, handleChange, touched, errors, isSubmitting }) => (
          <Form>
            <Stack spacing={3}>
              <TextField
                fullWidth
                label="ATP Set ID"
                name="atp_set_id"
                value={values.atp_set_id}
                onChange={handleChange}
                error={touched.atp_set_id && Boolean(errors.atp_set_id)}
                helperText={touched.atp_set_id && errors.atp_set_id}
              />

              <TextField
                fullWidth
                label="TP ID"
                name="tp_id"
                value={values.tp_id}
                onChange={handleChange}
                error={touched.tp_id && Boolean(errors.tp_id)}
                helperText={touched.tp_id && errors.tp_id}
              />

              <TextField
                fullWidth
                type="number"
                label="Sequence Number"
                name="sequence_number"
                value={values.sequence_number}
                onChange={handleChange}
                error={touched.sequence_number && Boolean(errors.sequence_number)}
                helperText={touched.sequence_number && errors.sequence_number}
              />

              <TextField
                fullWidth
                type="number"
                label="Week Number"
                name="week"
                value={values.week}
                onChange={handleChange}
                error={touched.week && Boolean(errors.week)}
                helperText={touched.week && errors.week}
              />

              <Stack direction="row" spacing={2} justifyContent="flex-end">
                {onCancel && (
                  <Button variant="outlined" onClick={onCancel} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Saving...' : isEdit ? 'Update ATP' : 'Create ATP'}
                </Button>
              </Stack>
            </Stack>
          </Form>
        )}
      </Formik>
    </Box>
  );
};
