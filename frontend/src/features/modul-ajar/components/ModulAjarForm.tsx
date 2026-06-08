/**
 * Modul Ajar Form Component
 * Form for creating and editing Modul Ajars
 */

import { Box, TextField, Button, Stack, Typography } from '@mui/material';
import { Formik, Form, FormikHelpers } from 'formik';
import * as Yup from 'yup';
import { CreateModulAjarRequest, UpdateModulAjarRequest } from '@/api/modul-ajar';

interface ModulAjarFormProps {
  initialValues?: Partial<CreateModulAjarRequest>;
  onSubmit: (values: CreateModulAjarRequest | UpdateModulAjarRequest) => void;
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
  const defaultValues: CreateModulAjarRequest = {
    modul_ajar_set_id: '',
    atp_id: '',
    tp_id: '',
    sequence_number: 1,
    title: '',
    learning_activities: '',
    teaching_methods: [],
    learning_media: [],
    learning_resources: [],
    attachments: [],
  };

  return (
    <Box>
      <Typography variant="h6" gutterBottom>
        {isEdit ? 'Edit Modul Ajar' : 'Create New Modul Ajar'}
      </Typography>
      <Formik
        initialValues={{ ...defaultValues, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values: CreateModulAjarRequest, helpers: FormikHelpers<CreateModulAjarRequest>) => {
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
