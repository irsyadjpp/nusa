/**
 * Modul Ajar Form Component
 * Form for creating and editing Modul Ajars
 */

import { Box, TextField, Button, Stack, Typography } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { ModulAjarCreateRequest, ModulAjarUpdateRequest } from '@/api/modul-ajar';

interface ModulAjarFormProps {
  initialValues?: Partial<ModulAjarCreateRequest>;
  onSubmit: (values: ModulAjarCreateRequest | ModulAjarUpdateRequest) => void;
  onCancel?: () => void;
  isEdit?: boolean;
}

const validationSchema = Yup.object({
  modul_ajar_set_id: Yup.string().required('Modul Ajar Set is required'),
  atp_id: Yup.string().required('ATP is required'),
  tp_id: Yup.string().required('TP is required'),
  sequence_number: Yup.number().required('Sequence number is required').min(1),
  title: Yup.string().required('Title is required'),
});

export const ModulAjarForm = ({ initialValues, onSubmit, onCancel, isEdit = false }: ModulAjarFormProps) => {
  const defaultValues: ModulAjarCreateRequest = {
    modul_ajar_set_id: '',
    atp_id: '',
    week: 1,
    session_number: 1,
    learning_objectives: [],
    learning_activities: {
      opening: [],
      core_activities: [],
      closing: [],
    },
    teaching_materials: {
      resources: [],
      media: [],
      references: [],
      core_materials: [],
      supporting_materials: [],
      digital_resources: [],
    },
    learning_methods: [],
    assessment_methods: [],
    time_allocation: {
      total_hours: 2,
      per_week_hours: 2,
      hours_per_week: 2,
      hours_per_session: 1,
      breakdown: {},
    },
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit Modul Ajar' : 'Create New Modul Ajar'}
      </Typography>
      <Formik
        initialValues={{ ...defaultValues, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values: ModulAjarCreateRequest, helpers: FormikHelpers<ModulAjarCreateRequest>) => {
          onSubmit(values);
          helpers.resetForm();
        }}
      >
        {({ values, handleChange, touched, errors, isSubmitting }) => (
          <Form>
            <Stack spacing={3}>
              <TextField
                fullWidth
                label="Modul Ajar Set ID"
                name="modul_ajar_set_id"
                value={values.modul_ajar_set_id}
                onChange={handleChange}
                error={touched.modul_ajar_set_id && Boolean(errors.modul_ajar_set_id)}
                helperText={touched.modul_ajar_set_id && errors.modul_ajar_set_id}
              />

              <TextField
                fullWidth
                label="ATP ID"
                name="atp_id"
                value={values.atp_id}
                onChange={handleChange}
                error={touched.atp_id && Boolean(errors.atp_id)}
                helperText={touched.atp_id && errors.atp_id}
              />

              <TextField
                fullWidth
                multiline
                rows={4}
                label="Learning Activities"
                name="learning_activities"
                value={values.learning_activities}
                onChange={handleChange}
              />

              <Stack direction="row" spacing={2} justifyContent="flex-end">
                {onCancel && (
                  <Button variant="outlined" onClick={onCancel} disabled={isSubmitting}>
                    Cancel
                  </Button>
                )}
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Saving...' : isEdit ? 'Update Modul Ajar' : 'Create Modul Ajar'}
                </Button>
              </Stack>
            </Stack>
          </Form>
        )}
      </Formik>
    </Box>
  );
};
