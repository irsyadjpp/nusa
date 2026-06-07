# 01_EDUCATION_DOMAIN_MODEL.md

## Foundation Document for Education Operating System Indonesia 2045

**Version**: 8.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT - LOCKED
**Alignment**: 100% aligned with 00A_NATIONAL_EDUCATION_DIRECTION_2045.md, 00B_PRODUCT_VISION.md, and 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md

**Purpose**: Define the complete business domain model for Education Operating System (NUSA), serving as the foundation for Domain-Driven Design (DDD), Event Storming, Bounded Context, Database Architecture, AI Agent Architecture, Product Roadmap, and MVP Development.

---

# SECTION 1 — Executive Summary

## Why This Domain Model Exists

The Education Operating System (NUSA) requires a foundational understanding of the education domain that transcends software implementation. This domain model captures the fundamental structures, relationships, and workflows of the Indonesian education ecosystem as defined by national policy, curriculum frameworks, and educational best practices.

## Education Operating System Paradigm

NUSA is fundamentally different from traditional School Information Systems (SIS). SIS focuses on administrative tasks. NUSA focuses on the core work of education: curriculum delivery, learning facilitation, assessment, student growth, and AI orchestration.

## Curriculum Before Software

Education is not a generic domain. Indonesian education has specific policy requirements, regulatory frameworks, curriculum structures, and cultural contexts. Software must serve curriculum, not impose its own structures. When software misrepresents the education domain, it fails the educational mission.

## AI-First Architecture

AI is not a feature addition but the core of the system. The target platform balance is 90% AI, 10% Human Validation. Every domain is designed with AI-First principles, where AI handles routine tasks and humans handle high-value work requiring judgment, empathy, and relationship building.

## Foundation Alignment

This domain model is fully aligned with:
- **00A_NATIONAL_EDUCATION_DIRECTION_2045.md**: The strategic direction and national vision
- **00B_PRODUCT_VISION.md**: The product vision for Education Operating System
- **00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**: The architectural principles and guardrails

All subsequent development must adhere to these foundation documents.

---

# SECTION 2 — National Education Architecture

## Eight-Layer Hierarchy

```
Indonesia Emas 2045
    ↓
National Human Capital Goals
    ↓
Graduate Profile (8 Dimensions)
    ↓
National Standards (SNP)
    ↓
Curriculum (CP, TP, ATP, Modul Ajar)
    ↓
Learning Process (Pembelajaran Mendalam)
    ↓
Assessment (Formative & Summative)
    ↓
School Improvement (SPMI & Quality Cycles)
```

### Layer 1: Indonesia Emas 2045
- **Purpose**: Ultimate national vision for becoming a developed, sovereign, advanced, and prosperous nation
- **Regulatory Basis**: UUD 1945, UU No. 20 Tahun 2003, RPJMN 2025-2045
- **Key Concepts**: Human capital development, digital sovereignty, global competitiveness, sustainable development

### Layer 2: National Human Capital Goals
- **Purpose**: Specific human capital outcomes required to achieve Indonesia Emas 2045
- **Regulatory Basis**: National Human Capital Strategy, Peta Jalan Pendidikan Indonesia 2025-2045
- **Key Concepts**: Quality workforce, innovation capacity, character development, lifelong learning

### Layer 3: Graduate Profile (8 Dimensions)
- **Purpose**: Characteristics and competencies Indonesian education should develop
- **Regulatory Basis**: Permendikdasmen No. 10 Tahun 2025
- **Key Concepts**: Eight dimensions (Keimanan & Ketakwaan, Kewargaan, Penalaran Kritis, Kreativitas, Kolaborasi, Kemandirian, Kesehatan, Komunikasi)
- **Definition**: Profil Lulusan 8 Dimensi defines the ultimate outcomes of Indonesian education

### Layer 4: National Standards (SNP)
- **Purpose**: Minimum quality criteria for all educational units
- **Regulatory Basis**: PP No. 57 Tahun 2021, PP No. 4 Tahun 2022, various Permendikdasmen
- **Key Concepts**: Eight interrelated standards (Isi, Kompetensi Lulusan, Proses, Penilaian, Pendidik, Sarana Prasarana, Pengelolaan, Pembiayaan)

### Layer 5: Curriculum
- **Purpose**: What students learn, how they learn, assessment of learning
- **Regulatory Basis**: Permendikdasmen No. 12 Tahun 2025, Keputusan BSKAP No. 046/H/KR/2025
- **Key Concepts**: Phase-based structure (Fase Fondasi, A-F), CP, TP, ATP, Modul Ajar, three components (Intrakurikuler, Kokurikuler, Ekstrakurikuler)

### Layer 6: Learning Process (Pembelajaran Mendalam)
- **Purpose**: How curriculum should be implemented
- **Regulatory Basis**: Panduan Pembelajaran dan Asesmen, Kajian Pembelajaran Mendalam
- **Key Concepts**: Understand-Apply-Reflect cycle, mindful-meaningful-joyful learning, Deep Learning pedagogy

### Layer 7: Assessment
- **Purpose**: Evaluate learning achievement and provide feedback
- **Regulatory Basis**: Panduan Pembelajaran dan Asesmen, National Assessment framework
- **Key Concepts**: Formative, summative, diagnostic, authentic, holistic assessment

### Layer 8: School Improvement
- **Purpose**: Continuous enhancement of educational quality
- **Regulatory Basis**: Naskah Akademik SPMI, Permendikbud No. 9 Tahun 2022
- **Key Concepts**: SPMI quality cycle (Mapping, Planning, Implementation, Evaluation, Control), Rapor Pendidikan

---

# SECTION 2.1 — Education Outcome Hierarchy

## Vertical Outcome Chain

The Education Outcome Hierarchy defines the complete vertical chain from national vision to daily learning activities. Every layer produces outcomes for the next layer, ensuring complete traceability from the highest national vision to the smallest learning activity.

```
Indonesia Emas 2045
        ↓
National Education Vision
        ↓
Human Capital Development
        ↓
Graduate Profile (8 Dimensions)
        ↓
National Standards
        ↓
Capaian Pembelajaran (CP)
        ↓
Tujuan Pembelajaran (TP)
        ↓
Alur Tujuan Pembelajaran (ATP)
        ↓
Modul Ajar
        ↓
Learning Activities
        ↓
Assessment Evidence
        ↓
Student Growth
        ↓
Graduate Outcomes
```

## Layer Descriptions

### Indonesia Emas 2045
- **Purpose**: Ultimate national vision for becoming a developed, sovereign, advanced, and prosperous nation
- **Outcome for Next Layer**: National Education Vision
- **Traceability**: All educational activities must contribute to Indonesia Emas 2045

### National Education Vision
- **Purpose**: Specific educational vision to achieve Indonesia Emas 2045
- **Outcome for Next Layer**: Human Capital Development Goals
- **Traceability**: All educational policies and practices must align with national education vision

### Human Capital Development
- **Purpose**: Specific human capital outcomes required for national development
- **Outcome for Next Layer**: Graduate Profile (8 Dimensions)
- **Traceability**: All educational outcomes must contribute to human capital development

### Graduate Profile (8 Dimensions)
- **Purpose**: Characteristics and competencies Indonesian education should develop
- **Outcome for Next Layer**: National Standards
- **Traceability**: All educational standards and curriculum must develop graduate profile dimensions

### National Standards
- **Purpose**: Minimum quality criteria for all educational units
- **Outcome for Next Layer**: Capaian Pembelajaran (CP)
- **Traceability**: All curriculum and learning must meet national standards

### Capaian Pembelajaran (CP)
- **Purpose**: Phase-based learning achievement expectations
- **Outcome for Next Layer**: Tujuan Pembelajaran (TP)
- **Traceability**: All learning objectives must derive from CP

### Tujuan Pembelajaran (TP)
- **Purpose**: Specific learning objectives for each phase
- **Outcome for Next Layer**: Alur Tujuan Pembelajaran (ATP)
- **Traceability**: All learning sequences must achieve TP

### Alur Tujuan Pembelajaran (ATP)
- **Purpose**: Sequenced learning objectives across phases
- **Outcome for Next Layer**: Modul Ajar
- **Traceability**: All learning modules must follow ATP sequences

### Modul Ajar
- **Purpose**: Detailed lesson plans and learning materials
- **Outcome for Next Layer**: Learning Activities
- **Traceability**: All learning activities must implement Modul Ajar

### Learning Activities
- **Purpose**: Actual learning experiences in classrooms
- **Outcome for Next Layer**: Assessment Evidence
- **Traceability**: All assessment must measure learning activities

### Assessment Evidence
- **Purpose**: Evidence of student learning and achievement
- **Outcome for Next Layer**: Student Growth
- **Traceability**: All growth data must come from assessment evidence

### Student Growth
- **Purpose**: Measurable progress in learning and development
- **Outcome for Next Layer**: Graduate Outcomes
- **Traceability**: All graduate outcomes must demonstrate student growth

### Graduate Outcomes
- **Purpose**: Final achievement of graduate profile dimensions
- **Outcome for Next Layer**: Human Capital Development
- **Traceability**: All graduate outcomes must contribute to human capital development

## Core Principles

### Every Learning Activity Must Contribute To Graduate Profile Achievement

No learning activity can stand alone. Every learning activity, assessment, and educational intervention must be traceable back to the Graduate Profile (8 Dimensions). If a learning activity cannot be traced to a specific graduate profile dimension, it should not exist in NUSA.

### Complete Traceability

Every data point, every assessment, every report, and every educational decision must be traceable through the entire hierarchy:
- From Indonesia Emas 2045 to daily learning activities
- From daily learning activities back to Indonesia Emas 2045
- No isolated activities without clear contribution to national outcomes

### Outcome-Driven Education

Education is outcome-driven, not activity-driven. Activities are means to achieve outcomes, not ends in themselves. The hierarchy ensures that:
- Every activity has a clear purpose
- Every purpose is aligned with higher-level outcomes
- Every outcome contributes to national vision

### Curriculum as Execution Layer

Curriculum (CP, TP, ATP, Modul Ajar) is the execution layer that operationalizes the Graduate Profile. Curriculum is not an end in itself but the means to achieve graduate profile outcomes.

## Strategic Implications

### For NUSA Architecture
- All data models must support traceability through the hierarchy
- All workflows must enforce alignment with graduate profile
- All AI agents must operate within this outcome framework

### For Teacher Practice
- Teachers must understand how daily activities contribute to graduate profile
- Teachers must be able to trace activities through the hierarchy
- Teachers must validate that activities contribute to outcomes

### For Assessment
- All assessments must measure progress toward graduate profile dimensions
- Assessment data must be traceable to learning objectives
- Assessment results must inform student growth toward graduate outcomes

### For Reporting
- All reports must communicate progress toward graduate profile dimensions
- Reports must show traceability from activities to outcomes
- Reports must support continuous improvement toward national vision

---

# SECTION 3 — Graduate Profile Domain

## Domain 1: Graduate Profile

**Purpose**: Ultimate characteristics and competencies Indonesian education seeks to develop through Profil Lulusan 8 Dimensi

**Regulatory Basis**: Permendikdasmen No. 10 Tahun 2025

**Key Concepts**: Eight dimensions with developmental phases, integrated development, ultimate goal of all educational activities, North Star for all domains

**Inputs**: Indonesia Emas 2045, National Human Capital Goals, cultural values, global competency frameworks, child development research

**Outputs**: Dimension definitions with indicators, assessment frameworks, curriculum design guidelines, teacher PD focus, reporting frameworks

**Relationships**: Upstream from national vision, downstream to all domains, lateral across all levels, root entity of entire system

**AI Responsibilities**: Dimension assessment, personalized development pathways, predictive modeling, natural language analysis, multimodal assessment, dimension tracking across learning activities

**Human Responsibilities**: Vision and values definition, dimension validation, contextual interpretation, ethical decisions, character formation guidance

**Human Validation Point**: All dimension definitions and assessment frameworks must be validated by national education standards experts

## Why It Matters

Profil Lulusan 8 Dimensi is the ROOT ENTITY of the entire Education Operating System. Every feature, workflow, AI agent, and data model must be traceable to these eight dimensions. If a feature cannot be traced to a specific graduate profile dimension, it should not exist in NUSA.

## Relationship to Every Domain

- **Curriculum**: CP, TP, ATP, and Modul Ajar are designed to achieve specific graduate profile dimensions
- **Learning**: Learning activities are designed to develop specific graduate profile dimensions
- **Assessment**: Assessment measures progress toward graduate profile dimensions
- **Reporting**: Reports communicate progress toward graduate profile dimensions
- **Teacher Growth**: Teacher competencies are aligned with the ability to develop graduate profile dimensions
- **School Improvement**: Quality metrics are defined in terms of graduate profile development
- **AI Personalization**: AI recommendations are based on their ability to contribute to graduate profile outcomes

## The Eight Dimensions

### 1. Keimanan dan Ketakwaan kepada Tuhan YME
Faith and piety toward God the Almighty
- Development: Religious understanding, moral character, spiritual practice
- Assessment: Behavioral observation, reflection, project evidence

### 2. Kewargaan
Citizenship
- Development: Civic knowledge, democratic participation, national identity
- Assessment: Civic engagement, community service, understanding of rights and responsibilities

### 3. Penalaran Kritis
Critical Reasoning
- Development: Analytical thinking, problem-solving, evidence-based reasoning
- Assessment: Problem-solving tasks, argumentation, critical analysis

### 4. Kreativitas
Creativity
- Development: Innovative thinking, creative expression, original solutions
- Assessment: Creative projects, innovation tasks, artistic expression

### 5. Kolaborasi
Collaboration
- Development: Teamwork, communication, conflict resolution
- Assessment: Group projects, peer collaboration, leadership

### 6. Kemandirian
Independence
- Development: Self-regulation, decision-making, responsibility
- Assessment: Self-directed tasks, independent projects, goal achievement

### 7. Kesehatan
Health
- Development: Physical health, mental health, healthy lifestyle
- Assessment: Health behaviors, physical fitness, mental wellbeing

### 8. Komunikasi
Communication
- Development: Effective expression, listening, multilingual communication
- Assessment: Presentations, written communication, multilingual tasks

---

# SECTION 4 — Strategic Domains

## Domain 2: Foundational Competencies

**Purpose**: Six essential abilities for PAUD-SD transition success

**Regulatory Basis**: Panduan Pemetaan Kemampuan Fondasi 2023, embedded in CP for Fase Fondasi and A

**Key Concepts**: Six competencies (Nilai Agama & Budi Pekerti, Kematangan Emosi, Keterampilan Sosial & Bahasa, Pemaknaan Belajar Positif, Keterampilan Motorik & Perawatan Diri, Kematangan Kognitif), bridge function, not CP/KD

**Inputs**: Child development research, early childhood best practices, cultural context, school readiness requirements

**Outputs**: Competency definitions, curriculum mapping, assessment approaches, teacher guidance, reporting frameworks

**Relationships**: Supports all subsequent learning, prerequisite for academic success, early expression of graduate profile

**AI Responsibilities**: Developmental screening, personalized support, early warning, parent guidance, progress tracking, competency gap identification

**Human Responsibilities**: Observation, parent communication, emotional support, developmental guidance, contextual interpretation

**Human Validation Point**: All competency assessments must be validated by teachers and parents

## Domain 3: Curriculum

**Purpose**: Define what students learn, how learning organized, learning outcomes

**Regulatory Basis**: Permendikdasmen No. 12 Tahun 2025, Keputusan BSKAP No. 046/H/KR/2025

**Key Concepts**: Phase-based structure (Fase Fondasi, A-F), CP, TP, ATP, Modul Ajar, three components (Intrakurikuler, Kokurikuler, Ekstrakurikuler), Kurikulum Satuan Pendidikan, muatan wajib

**Inputs**: Graduate profile, national standards, subject expertise, child development, local context, stakeholder input

**Outputs**: CP definitions, curriculum structure, subject documents, Kurikulum Satuan Pendidikan, resources, guidelines

**Relationships**: Upstream from graduate profile, downstream to learning planning, lateral alignment across phases

**AI Responsibilities**: CP alignment analysis, Kurikulum Satuan Pendidikan generation, cross-reference mapping, resource recommendation, gap analysis, personalized pathways, TP generation, ATP sequencing, Modul Ajar generation

**Human Responsibilities**: CP formulation, pedagogical judgment, local context adaptation, resource selection, quality validation

**Human Validation Point**: All curriculum materials must be reviewed and approved by teachers

## Curriculum Hierarchy Structure

The national curriculum (Kurikulum Nasional) follows a hierarchical structure that must be accurately represented in the system:

```
Subject (Mata Pelajaran)
    ↓
Phase (Fase)
    ↓
Element (Elemen)
    ↓
Subelement (Subelemen)
    ↓
Capaian Pembelajaran (CP)
```

### Subject (Mata Pelajaran)
- **Purpose**: Top-level curriculum organization by subject area
- **Examples**: Mathematics, Bahasa Indonesia, Science, Social Studies
- **Regulatory Basis**: National curriculum subject list

### Phase (Fase)
- **Purpose**: Educational phases that group learning by developmental stages
- **Examples**: Fase Fondasi, Fase A, Fase B, Fase C, Fase D, Fase E, Fase F
- **Regulatory Basis**: Kurikulum Merdeka phase structure

### Element (Elemen)
- **Purpose**: Major components within a phase for a subject
- **Examples**: Number and Operations, Algebra, Geometry, Statistics (for Mathematics)
- **Regulatory Basis**: Subject-specific curriculum structure

### Subelement (Subelemen)
- **Purpose**: Detailed breakdown of elements into specific learning areas
- **Examples**: Whole Numbers, Fractions, Decimals, Ratios (within Number and Operations)
- **Regulatory Basis**: Subject-specific curriculum structure

### Capaian Pembelajaran (CP)
- **Purpose**: Specific learning achievement expectations at the subelement level
- **Examples**: Students can add and subtract whole numbers up to 1000
- **Regulatory Basis**: CP definitions from national curriculum
- **Traceability**: Each CP must reference its subelement, element, phase, and subject

## Why This Hierarchy Matters

### Accurate Curriculum Representation
The flat model (Subject → Phase → CP) incorrectly represents the actual national curriculum structure. The hierarchical model (Subject → Phase → Element → Subelement → CP) accurately reflects how the national curriculum is organized.

### Curriculum Traceability
The hierarchy enables complete traceability from national standards to specific learning objectives:
- National standards define subjects and phases
- Subject-specific curriculum defines elements and subelements
- CP defines specific learning achievements at the subelement level

### AI Generation Context
AI agents require full curriculum hierarchy context to generate accurate and aligned learning materials:
- TP generation must understand the element and subelement context
- ATP sequencing must respect the hierarchical structure
- Modul Ajar generation must align with specific subelement objectives

### Reporting and Analytics
The hierarchy enables granular reporting and analytics:
- Track progress at subelement level
- Identify gaps at element level
- Report alignment with national curriculum at phase and subject level

## Implementation Requirements

### Database Schema
- Separate tables for subjects, phases, elements, subelements, and CP
- Foreign key relationships enforcing the hierarchy
- Indexes for efficient hierarchy traversal

### API Contract
- Endpoints for each hierarchy level (subjects, phases, elements, subelements, CP)
- Hierarchical query capabilities (e.g., get all CP for a subelement)
- Batch retrieval for full hierarchy context

### AI Prompt Inputs
- Full hierarchy context in AI generation requests
- Subject, phase, element, subelement information included
- Traceability metadata in generated outputs

### User Interface
- Hierarchical navigation for curriculum browsing
- Context-aware CP selection
- Visual hierarchy representation

## Domain 4: Learning Planning

**Purpose**: Translate curriculum into specific learning plans

**Regulatory Basis**: Panduan Pembelajaran dan Asesmen, Panduan Pengembangan Kurikulum Satuan Pendidikan

**Key Concepts**: Planning hierarchy (CP → TP → ATP → Modul Ajar), differentiation, collaborative planning

**Inputs**: CP, student data, resources, school context, time allocation, previous experiences

**Outputs**: TP, ATP, Modul Ajar, assessment plans, resource lists, differentiation strategies

**Relationships**: Based on CP, guides delivery, lateral alignment, integrated with assessment

**AI Responsibilities**: TP generation, ATP sequencing, Modul Ajar generation, differentiation suggestions, resource matching, time optimization, collaboration support

**Human Responsibilities**: Creative design, pedagogical decisions, relationship building, contextual adaptation, quality validation

**Human Validation Point**: All lesson plans must be reviewed and approved by teachers

## Domain 5: Learning Delivery

**Purpose**: Implement learning plans in classrooms and learning environments

**Regulatory Basis**: Panduan Pembelajaran dan Asesmen, Deep Learning framework, pedagogical guides

**Key Concepts**: Deep Learning implementation (mindful, meaningful, joyful), pedagogical practices, learning environment, partnerships, digital integration, real-time adaptation

**Inputs**: Modul Ajar, student readiness, resources, time constraints, classroom context

**Outputs**: Student engagement, learning artifacts, assessment evidence, teacher observations, adaptations

**Relationships**: Implements Modul Ajar, generates assessment evidence, coordinates across subjects

**AI Responsibilities**: Real-time assistance, engagement monitoring, content adaptation, question generation, misconception detection, resource retrieval, translation/accessibility, personalized learning pathways

**Human Responsibilities**: Teaching, emotional support, relationship building, pedagogical judgment, character formation, mentoring

**Human Validation Point**: All pedagogical decisions require human judgment

## Domain 6: Deep Learning Operating Layer

**Purpose**: Cross-domain pedagogical operating model for quality learning beyond surface-level understanding

**Regulatory Basis**: Kajian Akademik Pembelajaran Mendalam, embedded in Kurikulum Merdeka guidance

**Key Concepts**: Understand-Apply-Reflect cycle, mindful-meaningful-joyful learning, higher-order thinking, character integration, metacognitive awareness

**Inputs**: Curriculum content, student prior knowledge, learning context, resources, teacher expertise

**Outputs**: Deep understanding, transferable skills, character development, metacognitive awareness, motivation

**Relationships**: Informed by curriculum, assessed through authentic methods, consistent across subjects, operating layer for all learning domains

**AI Responsibilities**: Understanding assessment, application suggestions, reflection prompts, metacognition support, pathway optimization, analytics, personalized learning cycle optimization

**Human Responsibilities**: Philosophy, relationships, pedagogical judgment, contextual interpretation, character formation guidance

**Human Validation Point**: All pedagogical decisions require human judgment

### Deep Learning as Operating Layer

Deep Learning is not merely a domain but the operating model for all learning activities. Every learning experience must pass through the Understand-Apply-Reflect cycle:

```
Understand (Memahami)
    ↓
Apply (Menerapkan)
    ↓
Reflect (Merefleksikan)
```

This cycle applies across all learning domains: Literacy, Numeracy, Coding, AI Literacy, and all subject areas.

## Domain 7: Assessment

**Purpose**: Evaluate student learning, provide feedback, improve learning

**Regulatory Basis**: Panduan Pembelajaran dan Asesmen, Permendikdasmen regulations, National Assessment framework

**Key Concepts**: Formative, summative, diagnostic, authentic, holistic assessment; multiple methods; integration; feedback and reporting

**Inputs**: Learning objectives, student work, observations, self/peer assessments, standardized instruments

**Outputs**: Assessment results, feedback, progress reports, achievement records, diagnostic data, aggregate data

**Relationships**: Aligned with objectives, informs reporting and improvement, consistent practices, integrated with learning

**AI Responsibilities**: Automated scoring, feedback generation, pattern recognition, adaptive assessment, authentic evaluation, predictive analytics, multimodal assessment, rubric generation, remedial recommendations

**Human Responsibilities**: Complex evaluation, fairness judgment, contextualization, high-stakes decisions, ethical considerations

**Human Validation Point**: All assessment items must be validated for fairness; all high-stakes grading requires human approval

## Domain 8: Reporting

**Purpose**: Communicate student progress, achievement, development to stakeholders

**Regulatory Basis**: Panduan Pembelajaran dan Asesmen, Rapor Pendidikan framework, reporting requirements

**Key Concepts**: Stakeholders (students, parents, teachers, administrators, government), report types (progress, achievement, character, school), content, frequency, channels

**Inputs**: Assessment results, student work, observations, attendance records, teacher reflections

**Outputs**: Report cards, progress reports, character summaries, school reports, Rapor submissions, communication logs

**Relationships**: Based on assessment data, informs decisions, consistent formats, primary parent communication

**AI Responsibilities**: Report generation, narrative writing, translation, visualization, anomaly detection, recommendations, parent communication, data synthesis, progress communication

**Human Responsibilities**: Sensitive communication, contextualization, tone adjustment, decision interpretation, relationship building

**Human Validation Point**: All reports must be reviewed for accuracy and appropriateness before distribution

---

# SECTION 5 — Learning Domains

## Domain 9: Literacy

**Purpose**: Develop reading, writing, communication skills

**Regulatory Basis**: CP for Bahasa Indonesia, Gerakan Literasi Nasional, language education policies, National Assessment literacy

**Key Concepts**: Components (reading, writing, listening, speaking, viewing, digital), developmental progression, across curriculum, assessment, interventions

**Inputs**: Language development research, assessment data, student backgrounds, curriculum requirements, resources

**Outputs**: Assessment results, writing products, communication artifacts, progress data, intervention plans

**Relationships**: Based on research, enables learning across subjects, assessed specifically and integrated, supports numeracy

**AI Responsibilities**: Reading assessment, writing evaluation, personalized reading, speech recognition, literacy analytics, adaptive instruction, translation support

**Human Responsibilities**: Reading aloud, support, explanation, teaching, relationship building

**Human Validation Point**: All literacy assessments must be validated by teachers

## Domain 10: Numeracy

**Purpose**: Develop mathematical understanding, skills, dispositions

**Regulatory Basis**: CP for Matematika, Gerakan Numerasi Nasional 2026, mathematics policies, National Assessment numeracy

**Key Concepts**: Components (number sense, algebraic thinking, geometry, statistics, probability, communication, disposition), developmental progression, across curriculum, assessment, interventions

**Inputs**: Mathematics education research, assessment data, student backgrounds, curriculum requirements, resources

**Outputs**: Assessment results, mathematical solutions, reasoning artifacts, progress data, intervention plans

**Relationships**: Based on research, enables quantitative reasoning, assessed specifically and integrated, supported by literacy

**AI Responsibilities**: Problem-solving assessment, adaptive practice, error analysis, visual representation, real-world applications, numeracy analytics, communication assessment

**Human Responsibilities**: Explanation, teaching, relationship building, conceptual guidance

**Human Validation Point**: All numeracy assessments must be validated by teachers

## Domain 11: Coding

**Purpose**: Develop computational thinking and programming skills

**Regulatory Basis**: Kajian Coding & AI, Regulasi Implementasi, Panduan AI untuk Guru 2025, digital literacy policies

**Key Concepts**: Computational thinking (decomposition, pattern recognition, abstraction, algorithmic), coding skills, developmental progression, integration approaches, assessment

**Inputs**: CS education research, technology trends, student backgrounds, curriculum requirements, technology resources

**Outputs**: Computational thinking artifacts, programming projects, skill assessments, technology products

**Relationships**: Based on trends, prepares for digital future, integrates with subjects, related to STEM, part of future skills

**AI Responsibilities**: Code generation, debugging assistance, concept explanation, project suggestions, skill assessment, personalized learning, ethics guidance

**Human Responsibilities**: Ethics, concepts, teaching, relationship building, project guidance

**Human Validation Point**: All coding projects must be validated by teachers

## Domain 12: AI Literacy

**Purpose**: Develop AI understanding and responsible AI use

**Regulatory Basis**: Kajian Coding & AI, Regulasi Implementasi, Panduan AI untuk Guru 2025, digital literacy policies

**Key Concepts**: AI understanding, responsible AI use, AI ethics, prompt engineering, AI-assisted learning, developmental progression, integration approaches

**Inputs**: AI education research, technology trends, student backgrounds, curriculum requirements, technology resources

**Outputs**: AI demonstrations, AI literacy assessments, AI projects, responsible AI use evidence

**Relationships**: Based on trends, prepares for AI-native future, integrates with subjects, related to STEM, part of future skills

**AI Responsibilities**: Concept explanation, project suggestions, skill assessment, personalized learning, ethics guidance, AI tool recommendation

**Human Responsibilities**: Ethics, concepts, teaching, relationship building, responsible AI guidance

**Human Validation Point**: All AI literacy projects must be validated by teachers

---

# SECTION 6 — Human Development Domains

## Domain 13: Character Development

**Purpose**: Intentional development of values, attitudes, behaviors aligned with Profil Lulusan 8 Dimensi

**Regulatory Basis**: Permendikdasmen No. 10 Tahun 2025, Pendidikan Pancasila, Pendidikan Karakter, Antikorupsi 2026

**Key Concepts**: Dimensions aligned with Profil Lulusan 8 Dimensi, approaches (explicit, implicit, modeling, practice, community), integration points, specific programs, assessment

**Inputs**: National values, Pancasila philosophy, character research, family/community values, religious traditions

**Outputs**: Character evidence, behavioral observations, reflection journals, projects, assessment of growth

**Relationships**: Based on values, develops character, integrated across activities, directly supports graduate profile, requires home-school alignment

**AI Responsibilities**: Behavioral analysis, scenario simulation, reflection support, pattern recognition, analytics, situational judgment, recommendations

**Human Responsibilities**: Role modeling, relationship building, character formation guidance, contextual interpretation, ethical decisions

**Human Validation Point**: All character assessments must be validated by teachers and parents

## Domain 14: Student Wellbeing

**Purpose**: Support physical, mental, and emotional wellbeing of students

**Regulatory Basis**: Permendikdasmen regulations, child protection policies, mental health guidelines

**Key Concepts**: Physical health, mental health, emotional support, safe environment, wellbeing monitoring, intervention

**Inputs**: Health data, behavioral observations, student self-report, parent reports, school context

**Outputs**: Wellbeing assessments, intervention plans, support resources, progress tracking, referral documentation

**Relationships**: Supports all learning, prerequisite for academic success, integrated with character development, requires home-school alignment

**AI Responsibilities**: Wellbeing monitoring, early warning, pattern recognition, intervention recommendations, progress tracking

**Human Responsibilities**: Emotional support, relationship building, intervention decisions, contextual interpretation, ethical considerations

**Human Validation Point**: All wellbeing interventions must be validated by school counselors and parents

## Domain 15: Parent Partnership

**Purpose**: Engage parents and families in children's education

**Regulatory Basis**: Parent participation policies, communication requirements, family engagement best practices

**Key Concepts**: Partnership roles (supporters, decision-makers, advocates, learners, volunteers), communication channels, engagement activities, parent education, barriers (time, language, education, culture, economic, geographic)

**Inputs**: School policies, parent contact info, student data, events calendar, parent feedback, community resources

**Outputs**: Communication logs, meeting docs, volunteer coordination, parent education materials, engagement data, partnership agreements

**Relationships**: Based on policies, supports student success, lateral community building, primary reporting channel, critical for character, key to improvement

**AI Responsibilities**: Communication automation, translation, engagement analytics, resource recommendation, meeting scheduling, Q&A, personalization, home learning support

**Human Responsibilities**: Relationships, emotional support, cultural bridge, community building, decision-making

**Human Validation Point**: All parent communications must be reviewed by teachers

## Domain 16: Career & Future Readiness

**Purpose**: Prepare students for future education and career pathways

**Regulatory Basis**: Career guidance policies, future skills frameworks, industry partnership guidelines

**Key Concepts**: Career exploration, future skills, industry partnerships, pathway planning, readiness assessment

**Inputs**: Student interests, aptitude data, industry trends, labor market information, education pathways

**Outputs**: Career plans, readiness assessments, pathway recommendations, industry connections, skill development plans

**Relationships**: Informed by all learning domains, supports graduate profile dimensions, integrates with AI literacy and coding

**AI Responsibilities**: Career matching, pathway recommendation, skill gap analysis, industry trend analysis, personalized guidance

**Human Responsibilities**: Mentoring, relationship building, contextual interpretation, ethical guidance, decision support

**Human Validation Point**: All career guidance must be validated by school counselors

## Domain 17: Lifelong Learning Record

**Purpose**: Become the national lifelong learning record that connects all educational phases from early childhood through professional development

**Regulatory Basis**: National education policies, lifelong learning frameworks, digital credential standards, data protection regulations

**Key Concepts**: Continuous learning record, cross-phase connectivity, competency tracking, portfolio development, digital credentials, AI learning profile, career readiness, data portability, privacy-preserving analytics

**Inputs**: All learning activities, assessment results, achievements, certifications, projects, character evidence, wellbeing data, career milestones, professional development records

**Outputs**: Learning Passport, Digital Competency Transcript, Graduate Readiness Score, Career Readiness Profile, Portfolio, Achievement Records, AI Learning Profile

**Relationships**: Connects all educational phases, integrates with all domains, provides longitudinal view of student development, supports workforce planning

**AI Responsibilities**: Competency gap analysis, career recommendation, personalized learning path, workforce readiness prediction, learning pattern analysis, credential verification, portfolio curation

**Human Responsibilities**: Record validation, credential verification, privacy decisions, career guidance, contextual interpretation, ethical considerations

**Human Validation Point**: All credential verification and career recommendations must be validated by qualified professionals

## Lifelong Learning Phases

The Lifelong Learning Record Domain connects all educational phases:

```
PAUD (Early Childhood Education)
 ↓
SD (Primary School)
 ↓
SMP (Junior Secondary School)
 ↓
SMA/SMK (Senior Secondary School)
 ↓
Higher Education (University/Vocational)
 ↓
Professional Development
 ↓
Lifelong Learning
```

## Subdomains

### Student Profile Subdomain
- **Purpose**: Comprehensive student identity and demographic information
- **Components**: Personal information, educational history, family background, special needs, learning preferences
- **AI Opportunities**: Profile completion, learning preference analysis, special needs identification

### Competency Profile Subdomain
- **Purpose**: Track competency development across all phases
- **Components**: Subject competencies, graduate profile dimensions, foundational competencies, future skills
- **AI Opportunities**: Competency gap analysis, mastery prediction, personalized learning pathways

### Portfolio Subdomain
- **Purpose**: Curate and showcase student work and achievements
- **Components**: Projects, artifacts, creative work, research, presentations, performances
- **AI Opportunities**: Portfolio curation, achievement highlighting, portfolio recommendation

### Projects Subdomain
- **Purpose**: Track project-based learning experiences
- **Components**: Project descriptions, collaboration records, outcomes, reflections, assessments
- **AI Opportunities**: Project recommendation, skill mapping, collaboration analysis

### Achievements Subdomain
- **Purpose**: Record formal and informal achievements
- **Components**: Awards, competitions, certifications, recognitions, milestones
- **AI Opportunities**: Achievement matching, opportunity recommendation, skill validation

### Certificates Subdomain
- **Purpose**: Manage digital credentials and certificates
- **Components**: Course certificates, skill badges, professional certifications, micro-credentials
- **AI Opportunities**: Credential verification, skill validation, credential recommendation

### Character Evidence Subdomain
- **Purpose**: Document character development across all phases
- **Components**: Character observations, service activities, leadership experiences, community engagement
- **AI Opportunities**: Character pattern analysis, growth tracking, intervention recommendation

### AI Learning Profile Subdomain
- **Purpose**: Track AI-assisted learning patterns and preferences
- **Components**: AI interaction history, learning style analysis, AI tool usage, personalization effectiveness
- **AI Opportunities**: Learning optimization, personalization tuning, AI tool recommendation

### Career Readiness Profile Subdomain
- **Purpose**: Assess and track career readiness across all phases
- **Components**: Career interests, skill assessments, industry knowledge, work experiences, career goals
- **AI Opportunities**: Career matching, pathway recommendation, skill gap analysis, workforce readiness prediction

## Core Outputs

### Learning Passport
- **Description**: Comprehensive record of all learning experiences across phases
- **Components**: Educational history, competencies, achievements, certifications, projects
- **Use Cases**: College admissions, job applications, credential verification, lifelong learning planning

### Digital Competency Transcript
- **Description**: Machine-readable transcript of competencies and achievements
- **Components**: Competency levels, achievement dates, verification status, evidence links
- **Use Cases**: Automated credential verification, skill matching, workforce planning

### Graduate Readiness Score
- **Description**: Aggregate score measuring readiness for next educational phase or career
- **Components**: Academic readiness, character readiness, career readiness, wellbeing readiness
- **Use Cases**: Transition planning, intervention targeting, resource allocation

### Career Readiness Profile
- **Description**: Detailed profile of career-related skills and readiness
- **Components**: Industry-specific skills, soft skills, work experience, career interests
- **Use Cases**: Career guidance, job matching, workforce development

## Core Principles

### Learning Records Never Reset
Learning records continuously grow across a lifetime. They are never reset or deleted, only augmented with new experiences and achievements.

### Complete Traceability
Every learning experience, assessment, and achievement is traceable to its source and context. This ensures data integrity and enables verification.

### Privacy-Preserving Analytics
Analytics and insights are generated while preserving individual privacy. Data is aggregated and anonymized for analysis while maintaining individual records.

### Data Portability
Learning records are portable across institutions and systems. Students own their data and can transfer it between educational providers.

### Credential Verification
All credentials and achievements are verifiable through digital signatures and blockchain technology where appropriate.

## Future AI Opportunities

### Competency Gap Analysis
AI analyzes learning records to identify competency gaps and recommend targeted learning experiences to close gaps.

### Career Recommendation
AI matches learning records with career opportunities and recommends pathways based on skills, interests, and market trends.

### Personalized Learning Path
AI generates personalized learning pathways based on learning history, career goals, and competency gaps.

### Workforce Readiness Prediction
AI predicts workforce readiness based on learning records, industry trends, and skill requirements.

### Portfolio Curation
AI curates and highlights the most relevant achievements and projects for specific purposes (college applications, job applications, etc.).

## Strategic Implications

### For National Education Intelligence
- Provides longitudinal data on educational outcomes
- Enables evidence-based policy making
- Supports workforce planning and development
- Tracks progress toward Indonesia Emas 2045 goals

### For Student Empowerment
- Students own and control their learning records
- Enables informed educational and career decisions
- Supports lifelong learning and continuous development
- Provides verifiable credentials for opportunities

### For Educational Providers
- Enables seamless transitions between educational phases
- Supports personalized learning based on complete history
- Facilitates credential verification and recognition
- Improves educational quality through data-driven insights

### For Employers
- Provides verified competency and achievement records
- Enables skill-based hiring and matching
- Supports workforce development and training
- Reduces credential verification costs

---

# SECTION 7 — Teacher Domains

## Domain 18: Teacher Professional Growth

**Purpose**: Teacher development from initial preparation to ongoing learning

**Regulatory Basis**: Standar Pendidik, certification requirements, PD policies, performance assessment frameworks

**Key Concepts**: Development stages (pre-service, induction, novice, proficient, expert), content (subject matter, pedagogy, curriculum, assessment, digital, character, inclusive), modalities (formal, PLC, mentoring, action research, self-directed, online), assessment, career pathways

**Inputs**: Teacher standards, school needs, student data, curriculum changes, research, teacher self-assessment

**Outputs**: PD plans, certification records, performance results, learning artifacts, improvement evidence, career documentation

**Relationships**: Based on standards, improves teaching quality, lateral PLCs, requires curriculum understanding, drives school improvement

**AI Responsibilities**: Personalized learning, classroom analysis, resource matching, practice simulation, performance analytics, collaboration support, knowledge management, PD recommendation

**Human Responsibilities**: Mentoring, support, relationship building, professional judgment, contextual interpretation

**Human Validation Point**: All PD recommendations must be validated by teachers and school leaders

## Domain 19: Teacher Performance & Workload

**Purpose**: Monitor teacher performance and manage workload to prevent burnout

**Regulatory Basis**: Teacher workload policies, performance assessment frameworks, wellbeing guidelines

**Key Concepts**: Performance metrics, workload tracking, burnout prevention, support systems, recognition

**Inputs**: Performance data, workload data, self-assessment, peer feedback, student outcomes

**Outputs**: Performance reports, workload analyses, support recommendations, recognition documentation

**Relationships**: Informs teacher growth, supports school improvement, ensures sustainable teaching practice

**AI Responsibilities**: Performance analytics, workload monitoring, early warning for burnout, support recommendation, pattern recognition

**Human Responsibilities**: Performance evaluation, contextual interpretation, support decisions, relationship building

**Human Validation Point**: All performance evaluations must be validated by school leaders

---

# SECTION 8 — School Domains

## Domain 20: School Improvement

**Purpose**: Systematic process for schools to continuously enhance quality

**Regulatory Basis**: Naskah Akademik SPMI, Permendikbud No. 9 Tahun 2022, Rapor Pendidikan, accreditation

**Key Concepts**: SPMI, quality cycle (mapping, planning, implementation, evaluation, control), standards (SNP, kekhasan, benchmark, aspirational), data sources (Rapor, assessment, surveys, observations, admin), priorities, strategies

**Inputs**: National standards, Rapor data, assessment data, stakeholder feedback, best practices, school context

**Outputs**: Quality mapping, improvement plans, implementation docs, evaluation reports, quality data, accreditation reports

**Relationships**: Based on standards, improves all aspects, lateral learning, assessment data input, PD key strategy

**AI Responsibilities**: Quality analytics, benchmarking, predictive modeling, recommendation engine, progress tracking, resource optimization, best practice matching, school analytics

**Human Responsibilities**: Vision, community building, interpretation, decision-making, change management

**Human Validation Point**: All improvement strategies must be validated by school leaders

## Domain 21: Quality Assurance & Accreditation

**Purpose**: Ensure and validate educational quality through accreditation processes

**Regulatory Basis**: Accreditation standards, quality assurance frameworks, regulatory requirements

**Key Concepts**: Accreditation standards, quality assurance processes, compliance monitoring, external validation, continuous improvement

**Inputs**: School data, performance data, stakeholder feedback, compliance documentation

**Outputs**: Accreditation reports, quality assurance documentation, compliance status, improvement recommendations

**Relationships**: Informs school improvement, validates quality, ensures regulatory compliance

**AI Responsibilities**: Compliance monitoring, gap analysis, recommendation generation, predictive quality assessment

**Human Responsibilities**: External validation, interpretation, decision-making, relationship building with accrediting bodies

**Human Validation Point**: All accreditation decisions must be validated by accrediting bodies

---

# SECTION 9 — Platform Domains

## Domain 22: Education Data & Interoperability

**Purpose**: Ensure data integrity, interoperability, and single source of truth across the education system

**Regulatory Basis**: Data protection regulations, interoperability standards, data governance frameworks

**Key Concepts**: Single Source of Truth (SSOT), data hierarchy, data flow, interoperability, data governance, privacy, security

**Inputs**: All domain data, national standards, regulatory requirements

**Outputs**: Unified data architecture, interoperability standards, data governance documentation, data analytics foundation

**Relationships**: Foundational to all domains, enables data flow, supports AI operations, ensures national intelligence

**AI Responsibilities**: Data processing, analytics, pattern recognition, predictive modeling, data quality monitoring

**Human Responsibilities**: Data governance, privacy decisions, ethical considerations, regulatory compliance

**Human Validation Point**: All data governance decisions must be validated by data governance board

## Domain 23: AI Orchestration Domain

**Purpose**: Coordinate AI agents across all domains to provide seamless AI-native education experience

**Regulatory Basis**: AI ethics guidelines, data protection regulations, AI governance frameworks

**Key Concepts**: Multi-agent coordination, AI agent architecture, human-in-the-loop governance, AI ethics, AI transparency

**Inputs**: All domain data, user context, learning objectives, performance data

**Outputs**: AI agent coordination, AI recommendations, AI-generated content, AI analytics

**Relationships**: Cross-domain layer that supports all domains, enables AI-native experience, ensures human-in-the-loop governance

### Subdomain 22.1: Curriculum AI Agent

**Purpose**: Generate and optimize curriculum materials

**AI Responsibilities**: Generate TP, ATP, Modul Ajar, curriculum alignment analysis, resource recommendation

**Human Responsibilities**: Approve curriculum alignment, validate pedagogical quality, adapt to local context

**Human Validation Point**: All curriculum materials must be reviewed and approved by teachers

### Subdomain 22.2: Assessment AI Agent

**Purpose**: Generate and optimize assessment materials

**AI Responsibilities**: Generate rubrics, generate feedback, generate remedial activities, automated scoring, pattern recognition

**Human Responsibilities**: Validate fairness, contextualize results, make final competency decisions

**Human Validation Point**: All assessment items must be validated for fairness; all high-stakes grading requires human approval

### Subdomain 22.3: Reporting AI Agent

**Purpose**: Generate and optimize reporting materials

**AI Responsibilities**: Generate narrative reports, generate parent summaries, data synthesis, progress communication, translation

**Human Responsibilities**: Validate student growth story, contextualize findings, adjust tone, ensure appropriateness

**Human Validation Point**: All reports must be reviewed for accuracy and appropriateness before distribution

### Subdomain 22.4: Teacher Copilot

**Purpose**: Assist teachers in daily planning and professional activities

**AI Responsibilities**: Daily planning assistance, reflection support, resource search, classroom analysis, PD recommendation

**Human Responsibilities**: Make pedagogical decisions, build relationships, provide emotional support, validate recommendations

**Human Validation Point**: All recommendations must be validated by teachers

### Subdomain 22.5: Parent AI Agent

**Purpose**: Assist parents in supporting children's learning

**AI Responsibilities**: Home learning support, student growth monitoring, communication automation, resource recommendation, Q&A

**Human Responsibilities**: Build relationships, provide emotional support, make decisions about home learning, validate recommendations

**Human Validation Point**: All parent communications must be reviewed by teachers

### Subdomain 22.6: School AI Agent

**Purpose**: Assist school leaders in school improvement and management

**AI Responsibilities**: School analytics, school improvement planning, quality analytics, benchmarking, resource optimization

**Human Responsibilities**: Make decisions about improvement strategies, interpret analytics, lead change management, validate recommendations

**Human Validation Point**: All improvement strategies must be validated by school leaders

---

# SECTION 9.1 — Core vs Supporting Domain Classification

## Domain Classification Framework

All 22 domains in NUSA are classified into strategic categories based on their role in the education system. This classification guides prioritization, resource allocation, and architectural decisions.

## Core Domains

Domains that form the main engine of NUSA. These domains are essential for the system to function and deliver value to users.

### Graduate Profile Domain
- **Classification**: Core
- **Reason**: The North Star and root entity of the entire system. All other domains exist to develop the graduate profile dimensions.
- **Criticality**: Without this domain, the system has no purpose or direction.

### Curriculum Domain
- **Classification**: Core
- **Reason**: Defines what students learn and how learning is organized. The execution layer for graduate profile.
- **Criticality**: Without curriculum, there is no educational content to deliver.

### Learning Planning Domain
- **Classification**: Core
- **Reason**: Translates curriculum into specific learning plans. Core teacher work.
- **Criticality**: Without planning, learning delivery cannot happen effectively.

### Learning Delivery Domain
- **Classification**: Core
- **Reason**: Implements learning plans in classrooms. The actual teaching and learning process.
- **Criticality**: Without delivery, no learning occurs.

### Assessment Domain
- **Classification**: Core
- **Reason**: Evaluates student learning and provides feedback. Essential for measuring progress.
- **Criticality**: Without assessment, there is no way to know if learning is happening.

### Reporting Domain
- **Classification**: Core
- **Reason**: Communicates student progress to stakeholders. Primary parent communication channel.
- **Criticality**: Without reporting, stakeholders cannot understand progress.

## Strategic Domains

Domains that strengthen the quality of education. These domains enhance the core domains and drive continuous improvement.

### Deep Learning Operating Layer
- **Classification**: Strategic
- **Reason**: Cross-domain pedagogical operating model for quality learning beyond surface-level understanding.
- **Impact**: Elevates the quality of all learning activities.

### Character Development Domain
- **Classification**: Strategic
- **Reason**: Intentional development of values, attitudes, behaviors aligned with graduate profile.
- **Impact**: Directly supports multiple graduate profile dimensions.

### Teacher Professional Growth Domain
- **Classification**: Strategic
- **Reason**: Teacher development from initial preparation to ongoing learning.
- **Impact**: Improves the quality of all core domains through better teaching.

### School Improvement Domain
- **Classification**: Strategic
- **Reason**: Systematic process for schools to continuously enhance quality.
- **Impact**: Drives continuous improvement across all domains.

### Parent Partnership Domain
- **Classification**: Strategic
- **Reason**: Engages parents and families in children's education.
- **Impact**: Critical for student success and character development.

## Future Strategic Domains

The following domains are designated as **FUTURE STRATEGIC DOMAINS** and are explicitly excluded from MVP Wave 1 scope. These domains represent strategic capabilities that will be implemented in subsequent phases after MVP validation and foundation stabilization.

### Competency Graph Intelligence Domain
**Status**: FUTURE STRATEGIC DOMAIN
**Purpose**: Maintain and leverage the National Competency Graph to enable intelligent recommendations and personalization
**MVP Exclusion**: Not required for MVP Wave 1 curriculum-to-report workflow
**Implementation Phase**: Phase 2 or later

### Digital Twin Intelligence Domain
**Status**: FUTURE STRATEGIC DOMAIN
**Purpose**: Maintain digital representations of student learning journeys for personalization
**MVP Exclusion**: Not required for MVP Wave 1 curriculum-to-report workflow
**Implementation Phase**: Phase 2 or later

### Lifelong Learning Record Domain
**Status**: FUTURE STRATEGIC DOMAIN
**Purpose**: Maintain comprehensive longitudinal learning records across all educational phases
**MVP Exclusion**: Not required for MVP Wave 1 curriculum-to-report workflow
**Implementation Phase**: Phase 3 or later

### Education Analytics Domain
**Status**: FUTURE STRATEGIC DOMAIN
**Purpose**: Provide comprehensive education analytics and intelligence for decision-making
**MVP Exclusion**: Not required for MVP Wave 1 curriculum-to-report workflow
**Implementation Phase**: Phase 2 or later

### Quality Assurance & Accreditation Domain
**Status**: FUTURE STRATEGIC DOMAIN
**Purpose**: Support quality assurance processes and accreditation compliance
**MVP Exclusion**: Not required for MVP Wave 1 curriculum-to-report workflow
**Implementation Phase**: Phase 3 or later

## MVP Scope Protection

**MVP Wave 1 Domain Scope is Strictly Limited To**:
- Graduate Profile Domain
- Curriculum Domain
- Learning Planning Domain
- Learning Delivery Domain
- Assessment Domain
- Reporting Domain
- AI Orchestration Domain

**Explicitly Excluded from MVP Wave 1**:
- All domains listed in this Future Strategic Domains section
- All Strategic Domains listed above (Deep Learning, Character Development, Teacher Professional Growth, School Improvement, Parent Partnership)
- All Supporting Domains
- All Foundation Domains (except Foundational Competencies as needed for core operations)

No future strategic domain shall be included in MVP Wave 1 implementation without explicit architecture freeze amendment approved by Chief Enterprise Architect.

## Supporting Domains

Domains that enrich the learning experience. These domains provide specialized support and skills that enhance educational outcomes.

### Literacy Domain
- **Classification**: Supporting
- **Reason**: Develops reading, writing, communication skills. Enables learning across subjects.
- **Impact**: Foundational competency that supports all learning domains.

### Numeracy Domain
- **Classification**: Supporting
- **Reason**: Develops mathematical understanding, skills, dispositions. Enables quantitative reasoning.
- **Impact**: Foundational competency that supports STEM and analytical thinking.

### Coding Domain
- **Classification**: Supporting
- **Reason**: Develops computational thinking and programming skills. Prepares for digital future.
- **Impact**: Future skill that integrates with subjects and STEM.

### AI Literacy Domain
- **Classification**: Supporting
- **Reason**: Develops AI understanding and responsible AI use. Prepares for AI-native future.
- **Impact**: Future skill that enables AI-native education.

### Student Wellbeing Domain
- **Classification**: Supporting
- **Reason**: Supports physical, mental, and emotional wellbeing of students.
- **Impact**: Prerequisite for academic success and character development.

### Career & Future Readiness Domain
- **Classification**: Supporting
- **Reason**: Prepares students for future education and career pathways.
- **Impact**: Connects learning to future opportunities and workforce needs.

### Lifelong Learning Record Domain
- **Classification**: Strategic
- **Reason**: National lifelong learning record connecting all educational phases, strategic asset for workforce planning and lifelong learning ecosystem.
- **Impact**: Provides longitudinal view of student development, supports workforce planning, enables data portability, serves as strategic national asset.

## Foundation Domains

Domains that provide foundational support for early development and transitions.

### Foundational Competencies Domain
- **Classification**: Foundation
- **Reason**: Six essential abilities for PAUD-SD transition success.
- **Impact**: Bridge function that supports all subsequent learning.

### Teacher Performance & Workload Domain
- **Classification**: Foundation
- **Reason**: Monitors teacher performance and manages workload to prevent burnout.
- **Impact**: Ensures sustainable teaching practice and teacher wellbeing.

### Quality Assurance & Accreditation Domain
- **Classification**: Foundation
- **Reason**: Ensures and validates educational quality through accreditation processes.
- **Impact**: Validates quality and ensures regulatory compliance.

## Platform Domains

Domains that enable the system to operate. These domains provide the technical and data infrastructure.

### Education Data & Interoperability Domain
- **Classification**: Platform
- **Reason**: Ensures data integrity, interoperability, and single source of truth.
- **Impact**: Foundational to all domains, enables data flow, supports AI operations.

### AI Orchestration Domain
- **Classification**: Platform
- **Reason**: Coordinates AI agents across all domains to provide seamless AI-native education experience.
- **Impact**: Cross-domain layer that enables AI-native experience and human-in-the-loop governance.

## Classification Summary

| Category | Count | Purpose | Priority |
|----------|-------|---------|----------|
| **Core Domains** | 6 | Main engine of NUSA | Highest |
| **Strategic Domains** | 6 | Strengthen quality | High |
| **Supporting Domains** | 6 | Enrich experience | Medium |
| **Foundation Domains** | 3 | Foundational support | Medium |
| **Platform Domains** | 2 | Enable system operation | Highest |

## Strategic Implications

### For MVP Development
- **Phase 1**: Focus on Core Domains (Graduate Profile, Curriculum, Learning Planning, Assessment, Reporting)
- **Phase 2**: Add Strategic Domains (Deep Learning, Character Development, Teacher Growth, School Improvement, Parent Partnership, Lifelong Learning Record)
- **Phase 3**: Add Supporting Domains (Literacy, Numeracy, Coding, AI Literacy, Wellbeing, Career Readiness)
- **Phase 4**: Add Foundation and Platform Domains (Foundational Competencies, Teacher Performance, Quality Assurance, Data, AI Orchestration)

### For Resource Allocation
- Core Domains receive maximum investment and attention
- Strategic Domains receive significant investment for quality improvement
- Supporting Domains receive investment based on specific needs and priorities
- Foundation Domains receive investment to ensure sustainability
- Platform Domains receive investment to enable scale and reliability

### For Architecture
- Core Domains have the strictest data consistency requirements
- Strategic Domains have strong integration with Core Domains
- Supporting Domains have flexible integration patterns
- Foundation Domains have specialized data models and workflows
- Platform Domains have cross-cutting concerns and must support all other domains

---

# SECTION 10 — AI-Native Domain Architecture

## AI-Native Architecture

Every domain in NUSA is designed with AI-Native architecture, consisting of four layers:

### Human Layer
Human work that requires judgment, empathy, creativity, relationship building, and ethical decision-making. This layer cannot be automated and must remain human.

### AI Layer
AI work that handles routine tasks, provides recommendations, generates content, and processes data. This layer is designed to be 90% automated.

### Automation Layer
Fully automated workflows that execute without human intervention. This layer handles repetitive, rule-based tasks.

### Validation Layer
Human validation points where AI outputs are reviewed and approved before action. This layer ensures human-in-the-loop governance.

## AI-Native Domain Matrix

| Domain | Human Work | AI Work | Automation % | Human Validation |
|--------|-----------|---------|--------------|------------------|
| **Graduate Profile** | Vision, values, ethical decisions | Dimension research, indicator generation, assessment | 80% | Dimension definitions and frameworks |
| **Foundational Competencies** | Observation, parent communication | Screening analysis, progress tracking, early warning | 70% | Competency assessments |
| **Curriculum** | CP formulation, pedagogical judgment | TP generation, ATP sequencing, Modul Ajar generation, alignment analysis | 90% | All curriculum materials |
| **Learning Planning** | Creative design, relationships | TP generation, ATP sequencing, Modul Ajar generation, differentiation | 90% | All lesson plans |
| **Learning Delivery** | Teaching, emotional support, relationship building | Real-time assistance, engagement monitoring, content adaptation | 70% | All pedagogical decisions |
| **Deep Learning** | Philosophy, relationships, pedagogical judgment | Understanding assessment, application suggestions, reflection prompts | 80% | All pedagogical decisions |
| **Assessment** | Complex evaluation, fairness judgment | Automated scoring, feedback generation, pattern recognition, rubric generation | 80% | All assessment items and high-stakes grading |
| **Reporting** | Sensitive communication, contextualization | Report generation, narrative writing, translation, data synthesis | 95% | All reports |
| **Literacy** | Reading aloud, support, explanation | Reading assessment, writing evaluation, personalized reading | 80% | All literacy assessments |
| **Numeracy** | Explanation, teaching, conceptual guidance | Problem-solving assessment, adaptive practice, error analysis | 80% | All numeracy assessments |
| **Coding** | Ethics, concepts, teaching | Code generation, debugging assistance, concept explanation | 85% | All coding projects |
| **AI Literacy** | Ethics, concepts, teaching | Concept explanation, project suggestions, skill assessment | 85% | All AI literacy projects |
| **Character Development** | Role modeling, relationship building | Behavioral analysis, scenario simulation, reflection support | 70% | All character assessments |
| **Student Wellbeing** | Emotional support, intervention decisions | Wellbeing monitoring, early warning, pattern recognition | 75% | All wellbeing interventions |
| **Parent Partnership** | Relationships, emotional support | Communication automation, translation, engagement analytics | 85% | All parent communications |
| **Career & Future Readiness** | Mentoring, relationship building | Career matching, pathway recommendation, skill gap analysis | 80% | All career guidance |
| **Lifelong Learning Record** | Record validation, credential verification, privacy decisions | Competency gap analysis, career recommendation, personalized learning path, workforce readiness prediction | 85% | All credential verification and career recommendations |
| **Teacher Professional Growth** | Mentoring, support, relationship building | Personalized learning, classroom analysis, resource matching | 85% | All PD recommendations |
| **Teacher Performance & Workload** | Performance evaluation, support decisions | Performance analytics, workload monitoring, early warning | 80% | All performance evaluations |
| **School Improvement** | Vision, community building, interpretation | Quality analytics, benchmarking, predictive modeling | 85% | All improvement strategies |
| **Quality Assurance & Accreditation** | External validation, interpretation | Compliance monitoring, gap analysis, recommendation generation | 80% | All accreditation decisions |
| **Education Data & Interoperability** | Data governance, privacy decisions | Data processing, analytics, pattern recognition | 90% | All data governance decisions |
| **AI Orchestration** | AI governance, ethical decisions | Multi-agent coordination, AI recommendations, AI-generated content | 95% | All AI governance decisions |

## Target Balance

The long-term target for NUSA is:
- **90% Automation**: Routine tasks handled by AI
- **10% Human Judgment**: High-value decisions requiring human judgment

This balance ensures that AI augments human capacity without replacing human judgment, empathy, and relationship building.

---

# SECTION 11 — Education Data Architecture

## Single Source of Truth (SSOT)

NUSA operates on the principle of Single Source of Truth (SSOT). Educational data must be entered once at the point of origin and automatically flow through the system. Data must not be duplicated, fragmented, or manually reconciled across multiple systems.

## SSOT Hierarchy

```
Indonesia Emas 2045
    ↓
Graduate Profile (8 Dimensions)
    ↓
National Standards (SNP)
    ↓
Capaian Pembelajaran (CP)
    ↓
Tujuan Pembelajaran (TP)
    ↓
Alur Tujuan Pembelajaran (ATP)
    ↓
Modul Ajar
    ↓
Learning Activities
    ↓
Assessment Evidence
    ↓
Student Portfolio
    ↓
Reporting
    ↓
School Improvement
```

## SSOT Principles

### No Data Duplication
- Educational data is not duplicated across multiple systems
- Data is entered once at the point of origin
- Data flows automatically through the system

### No Data Fragmentation
- Educational data is not fragmented across silos
- Data is unified in a single data architecture
- Data is accessible across all domains

### No Manual Reconciliation
- Educational data does not require manual reconciliation
- Data consistency is maintained automatically
- Data integrity is ensured through architecture

### Complete Traceability
- Every data point can be traced to its origin
- Every data point can be traced to curriculum objectives
- Every data point can be traced to graduate profile dimensions

## SSOT Benefits

### Data Integrity
- Consistent data across all uses
- No inconsistent records
- Complete audit trail
- Evidence-based decisions

### Efficiency
- No duplicate entry
- Automated data flow
- Reduced administrative burden
- Increased teacher productivity

### Scalability
- National aggregation possible
- Standardized data across all schools
- Comparison capability
- National education intelligence

---

# SECTION 12 — Domain Dependency Architecture

## Primary Dependency Chain

```
Indonesia Emas 2045
    ↓
National Human Capital Goals
    ↓
Graduate Profile (8 Dimensions)
    ↓
Foundational Competencies
    ↓
Curriculum (CP, TP, ATP, Modul Ajar)
    ↓
Learning Planning
    ↓
Learning Delivery
    ↓
Assessment
    ↓
Reporting
```

## Supporting Domains

### Literacy
- Supports all learning domains
- Enables learning across subjects
- Assessed specifically and integrated

### Numeracy
- Enables quantitative reasoning
- Assessed specifically and integrated
- Supported by literacy

### Coding
- Prepares for digital future
- Integrates with subjects
- Related to STEM

### AI Literacy
- Prepares for AI-native future
- Integrates with subjects
- Related to STEM

## Horizontal Domains

### Parent Partnership
- Supports student success
- Lateral community building
- Primary reporting channel
- Critical for character development

### Teacher Growth
- Improves teaching quality
- Lateral PLCs
- Requires curriculum understanding
- Drives school improvement

### Student Wellbeing
- Supports all learning
- Prerequisite for academic success
- Integrated with character development
- Requires home-school alignment

### Career & Future Readiness
- Informed by all learning domains
- Supports graduate profile dimensions
- Integrates with AI literacy and coding

## AI Orchestration Module

The AI Orchestration Module is a cross-domain layer that supports all domains:
- Coordinates AI agents across all domains
- Enables AI-native experience
- Ensures human-in-the-loop governance
- Provides seamless AI assistance

---

# SECTION 13 — Education Flywheel

## Core Flywheel

The Education Flywheel is the core process that the entire platform must support. The flywheel creates a virtuous cycle of continuous improvement:

```
Indonesia Emas 2045
    ↓
Graduate Profile (8 Dimensions)
    ↓
Curriculum (CP, TP, ATP, Modul Ajar)
    ↓
Learning Planning
    ↓
Learning Delivery
    ↓
Assessment
    ↓
Reporting
    ↓
Teacher Growth
    ↓
School Improvement
    ↓
Curriculum Improvement
    ↺ (back to Curriculum)
```

## Flywheel Stages

### Indonesia Emas 2045
**Input**: National vision for becoming a developed nation
**Output**: National Human Capital Goals
**Purpose**: Provides the ultimate context for all educational activities

### Graduate Profile (8 Dimensions)
**Input**: National Human Capital Goals
**Output**: Dimension definitions and assessment frameworks
**Purpose**: Defines the ultimate outcomes of education

### Curriculum
**Input**: Graduate Profile dimensions, National Standards
**Output**: CP, TP, ATP, Modul Ajar
**Purpose**: Operationalizes graduate profile through curriculum structures

### Learning Planning
**Input**: CP, TP, ATP, Modul Ajar
**Output**: Lesson plans, learning resources, differentiation strategies
**Purpose**: Prepares for effective learning delivery

### Learning Delivery
**Input**: Lesson plans, learning resources, differentiation strategies
**Output**: Learning activities, student engagement, learning progress
**Purpose**: Delivers learning experiences that achieve curriculum objectives

### Assessment
**Input**: Learning activities, learning progress
**Output**: Assessment evidence, learning gaps, graduate profile progress
**Purpose**: Measures progress toward curriculum objectives and graduate profile development

### Reporting
**Input**: Assessment evidence, learning gaps, graduate profile progress
**Output**: Progress narratives, quality metrics, improvement insights
**Purpose**: Communicates progress and generates insights for improvement

### Teacher Growth
**Input**: Progress narratives, quality metrics, improvement insights
**Output**: Teacher development plans, performance improvements
**Purpose**: Improves teacher capacity to deliver quality education

### School Improvement
**Input**: Teacher development, performance data, quality metrics
**Output**: Improvement strategies, curriculum refinements
**Purpose**: Drives continuous improvement at the school level

### Curriculum Improvement
**Input**: Improvement strategies, curriculum refinements
**Output**: Updated CP, TP, ATP, Modul Ajar
**Purpose**: Refines curriculum based on evidence and insights

## Continuous Data Flow

The flywheel ensures that data flows continuously through the system:
- Assessment data informs reporting
- Reporting data informs teacher growth
- Teacher growth data informs school improvement
- School improvement data informs curriculum refinement
- Curriculum refinement improves learning delivery
- Improved learning delivery generates better assessment data

This continuous flow creates a virtuous cycle of continuous improvement.

---

# SECTION 14 — AI Companion Mapping

## North Star Goal

From 00B_PRODUCT_VISION.md:

> "Every Indonesian Learner Has a Personal AI Learning Companion"

## AI Companion Role by Domain

| Domain | AI Companion Role |
|--------|------------------|
| **Graduate Profile** | Personal development coach tracking progress across 8 dimensions |
| **Foundational Competencies** | Early childhood development guide for parents and teachers |
| **Curriculum** | Curriculum navigator helping teachers navigate CP, TP, ATP, Modul Ajar |
| **Learning Planning** | Planning assistant generating lesson plans and resources |
| **Learning Delivery** | Real-time learning coach providing support during lessons |
| **Deep Learning** | Pedagogy coach guiding Understand-Apply-Reflect cycle |
| **Assessment** | Assessment assistant generating rubrics, feedback, and remedial activities |
| **Reporting** | Communication assistant generating narratives and parent summaries |
| **Literacy** | Reading and writing coach providing personalized literacy support |
| **Numeracy** | Math coach providing personalized numeracy support |
| **Coding** | Coding tutor providing code generation and debugging assistance |
| **AI Literacy** | AI guide teaching responsible AI use and prompt engineering |
| **Character Development** | Character coach supporting values and character formation |
| **Student Wellbeing** | Wellbeing monitor tracking physical and mental health |
| **Parent Partnership** | Parent guide supporting home learning and engagement |
| **Career & Future Readiness** | Career counselor providing pathway guidance |
| **Teacher Professional Growth** | PD coach recommending professional development |
| **Teacher Performance & Workload** | Workload advisor monitoring burnout risk |
| **School Improvement** | Improvement analyst providing quality analytics |
| **Quality Assurance & Accreditation** | Compliance advisor monitoring regulatory requirements |
| **Education Data & Interoperability** | Data steward ensuring data integrity |
| **AI Orchestration** | AI coordinator managing all AI agents |

## Unified AI Companion Experience

The AI Companion is not a single agent but a unified experience that coordinates across all domains:
- The AI Companion adapts to the user's role (student, teacher, parent, school leader)
- The AI Companion provides context-aware support based on the current domain
- The AI Companion maintains consistency across all interactions
- The AI Companion learns from user behavior to provide personalized support

---

# SECTION 15 — Platform Layer Mapping

## Platform Layer Architecture

From 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md, NUSA is architected as a layered platform where each layer provides specific capabilities and depends on the integrity of the layers below.

## Domain-to-Layer Mapping

### Layer 1: National Education Vision Layer
**Domains**: None (this layer encodes the vision, not domains)
**Purpose**: Encode Indonesia Emas 2045 and national human capital goals

### Layer 2: Curriculum Layer
**Domains**: Graduate Profile, Foundational Competencies, Curriculum
**Purpose**: Operationalize the graduate profile through CP, TP, ATP, and Modul Ajar

### Layer 3: Learning Layer
**Domains**: Learning Planning, Learning Delivery, Deep Learning, Literacy, Numeracy, Coding, AI Literacy
**Purpose**: Deliver curriculum through personalized learning pathways and Deep Learning pedagogy

### Layer 4: Assessment Layer
**Domains**: Assessment
**Purpose**: Measure learning progress through formative and summative assessment

### Layer 5: AI Layer
**Domains**: AI Orchestration, AI Companion (cross-domain)
**Purpose**: Process educational data to provide insights, predictions, and recommendations

### Layer 6: Data Layer
**Domains**: Education Data & Interoperability
**Purpose**: Aggregate and analyze educational data to provide insights for improvement

### Layer 7: Experience Layer
**Domains**: All domains (user interface and experience)
**Purpose**: Provide the user interface and experience for all stakeholders

## Cross-Layer Domains

### Human Development Domains
**Domains**: Character Development, Student Wellbeing, Parent Partnership, Career & Future Readiness
**Layers**: Experience Layer (primary), Data Layer (supporting)

### Teacher Domains
**Domains**: Teacher Professional Growth, Teacher Performance & Workload
**Layers**: Experience Layer (primary), Data Layer (supporting)

### School Domains
**Domains**: School Improvement, Quality Assurance & Accreditation
**Layers**: Experience Layer (primary), Data Layer (supporting)

---

# SECTION 15.1 — Bounded Context Map

## Domain-Driven Design (DDD) Bounded Context Map

The Bounded Context Map defines how domains interact through Domain-Driven Design patterns. Each bounded context represents a specific responsibility boundary with its own ubiquitous language.

## Primary Context Chain

```
Graduate Profile Context
        ↓

Curriculum Context
        ↓

Learning Context
        ↓

Assessment Context
        ↓

Reporting Context
```

## Context Relationships

### Graduate Profile Context
- **Role**: Upstream Context (Source of Truth)
- **Responsibility**: Defines the ultimate outcomes and dimensions
- **Relationships**: 
  - Downstream to: Curriculum Context, Learning Context, Assessment Context, Reporting Context
  - Shared Kernel: Dimension definitions, competency indicators
  - Pattern: Upstream/Downstream - All other contexts depend on Graduate Profile definitions

### Curriculum Context
- **Role**: Upstream Context (for Learning)
- **Responsibility**: Defines what students learn and how learning is organized
- **Relationships**:
  - Upstream from: Graduate Profile Context
  - Downstream to: Learning Context, Assessment Context
  - Shared Kernel: CP, TP, ATP, Modul Ajar structures
  - Pattern: Upstream/Downstream - Learning depends on Curriculum

### Learning Context
- **Role**: Downstream Context (from Curriculum)
- **Responsibility**: Implements curriculum through learning activities
- **Relationships**:
  - Upstream from: Curriculum Context, Graduate Profile Context
  - Downstream to: Assessment Context
  - Shared Kernel: Learning activity structures, student data
  - Pattern: Customer-Supplier - Learning consumes Curriculum, provides evidence to Assessment

### Assessment Context
- **Role**: Downstream Context (from Learning)
- **Responsibility**: Evaluates learning and provides feedback
- **Relationships**:
  - Upstream from: Learning Context, Curriculum Context, Graduate Profile Context
  - Downstream to: Reporting Context
  - Shared Kernel: Assessment structures, rubrics, evidence
  - Pattern: Customer-Supplier - Assessment consumes Learning evidence, provides data to Reporting

### Reporting Context
- **Role**: Downstream Context (from Assessment)
- **Responsibility**: Communicates progress to stakeholders
- **Relationships**:
  - Upstream from: Assessment Context, Learning Context, Curriculum Context, Graduate Profile Context
  - Shared Kernel: Report structures, student progress data
  - Pattern: Conformist - Reporting follows the data structures from upstream contexts

## Supporting Contexts

### Teacher Growth Context
- **Role**: Horizontal Context
- **Responsibility**: Supports teacher development across all domains
- **Relationships**:
  - Shared Kernel: Teacher data, performance metrics
  - Pattern: Open Host Service - Provides teacher development services to all contexts

### School Improvement Context
- **Role**: Horizontal Context
- **Responsibility**: Drives continuous improvement at school level
- **Relationships**:
  - Shared Kernel: Quality metrics, improvement data
  - Pattern: Open Host Service - Provides quality analytics to all contexts

### Parent Partnership Context
- **Role**: Horizontal Context
- **Responsibility**: Engages parents in education
- **Relationships**:
  - Shared Kernel: Parent data, communication logs
  - Pattern: Open Host Service - Provides communication services to all contexts

## Context Relationship Patterns

### Upstream Context
**Definition**: The source of truth for specific data or concepts. The upstream context defines the ubiquitous language and data structures that downstream contexts must follow.

**Examples**:
- Graduate Profile Context is upstream for all contexts
- Curriculum Context is upstream for Learning and Assessment contexts

**Implications**:
- Downstream contexts cannot change upstream definitions
- Changes in upstream contexts require coordination with downstream contexts
- Upstream contexts have authority over their domain language

### Downstream Context
**Definition**: Consumes data and concepts from upstream contexts. Must conform to the definitions and structures provided by upstream contexts.

**Examples**:
- Learning Context is downstream from Curriculum Context
- Assessment Context is downstream from Learning Context
- Reporting Context is downstream from Assessment Context

**Implications**:
- Must adapt to upstream changes
- Cannot redefine upstream concepts
- May add local extensions but must maintain compatibility

### Shared Kernel
**Definition**: A model or set of classes that is shared between multiple bounded contexts. The shared kernel represents the common language and data structures that contexts agree on.

**Examples**:
- Dimension definitions shared between Graduate Profile, Curriculum, Learning, Assessment, and Reporting
- Student data shared across all contexts
- Teacher data shared across Teacher Growth and School Improvement contexts

**Implications**:
- Changes to shared kernel require coordination across all contexts
- Shared kernel must be stable and well-governed
- Minimize shared kernel to reduce coupling

### Customer-Supplier
**Definition**: A relationship where one context (customer) depends on another context (supplier) for specific modules or data. The supplier provides value to the customer.

**Examples**:
- Learning Context (customer) depends on Curriculum Context (supplier) for learning materials
- Assessment Context (customer) depends on Learning Context (supplier) for evidence
- Reporting Context (customer) depends on Assessment Context (supplier) for data

**Implications**:
- Supplier must understand customer needs
- Customer must communicate requirements clearly
- Both contexts must agree on service contracts

### Conformist
**Definition**: A relationship where one context adopts the model and language of another context without having influence over it. The conformist context follows the rules of the dominant context.

**Examples**:
- Reporting Context conforms to Assessment Context data structures
- All contexts conform to Graduate Profile Context definitions

**Implications**:
- Conformist context has limited influence
- Must adapt to upstream changes
- Reduces coordination overhead but increases dependency

### Open Host Service
**Definition**: A relationship where a context provides a protocol or service that other contexts can use. The host context defines the service interface.

**Examples**:
- Teacher Growth Context provides PD services to all contexts
- School Improvement Context provides analytics services to all contexts
- Parent Partnership Context provides communication services to all contexts

**Implications**:
- Host context must maintain stable service interface
- Consuming contexts can integrate easily
- Enables horizontal integration without tight coupling

## Cross-Context Integration

### AI Orchestration Context
- **Role**: Cross-Context Service
- **Responsibility**: Coordinates AI agents across all contexts
- **Relationships**:
  - Pattern: Open Host Service - Provides AI services to all contexts
  - Integration: AI agents operate within each context but are coordinated centrally

### Education Data Context
- **Role**: Cross-Context Infrastructure
- **Responsibility**: Provides data infrastructure and SSOT across all contexts
- **Relationships**:
  - Pattern: Shared Kernel - Data structures and governance shared across all contexts
  - Integration: All contexts depend on data infrastructure for SSOT

## Strategic Implications

### For System Architecture
- Context boundaries define microservice boundaries
- Shared kernel defines shared libraries and data models
- Upstream/Downstream relationships define API contracts
- Open Host Service defines integration patterns

### For Data Architecture
- SSOT is maintained through shared kernel
- Data flow follows upstream/downstream relationships
- Data governance follows context authority
- Cross-context data access follows service patterns

### For AI Architecture
- AI agents operate within specific contexts
- AI Orchestration Context coordinates across contexts
- AI services follow Open Host Service pattern
- AI data follows SSOT principles

### For Development Teams
- Each bounded context can be developed by a dedicated team
- Context boundaries define team responsibilities
- Shared kernel requires cross-team coordination
- API contracts define inter-team communication

---

# SECTION 15.2 — Aggregate Root Definitions

## Domain-Driven Design Aggregate Roots

Aggregate Roots are the core entities in Domain-Driven Design that maintain consistency within their boundaries. Each Aggregate Root is the single entry point for accessing and modifying the entities within its aggregate.

## Graduate Profile Aggregate

```
Graduate Profile (Aggregate Root)
 ├── Dimension (Entity)
 │   ├── Dimension ID
 │   ├── Dimension Name
 │   ├── Dimension Description
 │   └── Developmental Phases
 ├── Competency (Entity)
 │   ├── Competency ID
 │   ├── Competency Name
 │   ├── Competency Description
 │   └── Dimension Reference
 └── Indicator (Entity)
     ├── Indicator ID
     ├── Indicator Description
     ├── Assessment Criteria
     └── Competency Reference
```

**Why This Aggregate Root Matters**:
- The Graduate Profile is the root entity of the entire NUSA system
- All other domains derive their purpose from this aggregate
- Ensures consistency in how graduate profile dimensions are defined and used
- Provides single source of truth for all competency definitions
- Enables traceability from daily activities to national outcomes

**Invariants**:
- All dimensions must align with Permendikdasmen No. 10 Tahun 2025
- All competencies must reference a valid dimension
- All indicators must reference a valid competency
- Dimension definitions cannot be changed without national approval

## Curriculum Aggregate

```
Curriculum (Aggregate Root)
 ├── CP (Capaian Pembelajaran) (Entity)
 │   ├── CP ID
 │   ├── Phase
 │   ├── Subject
 │   ├── Learning Outcomes
 │   └── Dimension References
 ├── TP (Tujuan Pembelajaran) (Entity)
 │   ├── TP ID
 │   ├── Phase
 │   ├── Subject
 │   ├── Learning Objectives
 │   └── CP Reference
 ├── ATP (Alur Tujuan Pembelajaran) (Entity)
 │   ├── ATP ID
 │   ├── Phase
 │   ├── Subject
 │   ├── Learning Sequence
 │   └── TP References
 └── Learning Unit (Entity)
     ├── Unit ID
     ├── Unit Name
     ├── Duration
     ├── TP References
     └── Resources
```

**Why This Aggregate Root Matters**:
- Curriculum is the execution layer for the graduate profile
- Ensures consistency in how curriculum is structured and delivered
- Provides single source of truth for all curriculum materials
- Enables traceability from learning objectives to graduate profile dimensions
- Supports Kurikulum Satuan Pendidikan generation

**Invariants**:
- All TP must reference a valid CP
- All ATP must reference valid TP sequences
- All Learning Units must reference valid ATP
- Curriculum must align with national CP and graduate profile dimensions

## Learning Aggregate

```
Learning Experience (Aggregate Root)
 ├── Activity (Entity)
 │   ├── Activity ID
 │   ├── Activity Type
 │   ├── Learning Objectives
 │   ├── Resources
 │   └── Modul Ajar Reference
 ├── Resource (Entity)
 │   ├── Resource ID
 │   ├── Resource Type
 │   ├── Resource Content
 │   └── Activity Reference
 └── Reflection (Entity)
     ├── Reflection ID
     ├── Student Reflection
     ├── Teacher Reflection
     ├── Learning Outcomes
     └── Activity Reference
```

**Why This Aggregate Root Matters**:
- Learning Experience is the actual implementation of curriculum
- Ensures consistency in how learning activities are designed and delivered
- Provides single source of truth for all learning data
- Enables traceability from activities to learning objectives
- Supports Deep Learning pedagogy implementation

**Invariants**:
- All Activities must reference a valid Modul Ajar
- All Resources must reference a valid Activity
- All Reflections must reference a valid Activity
- Learning activities must align with Deep Learning principles

## Assessment Aggregate

```
Assessment (Aggregate Root)
 ├── Evidence (Entity)
 │   ├── Evidence ID
 │   ├── Evidence Type
 │   ├── Student Work
 │   ├── Learning Objectives
 │   └── Activity Reference
 ├── Rubric (Entity)
 │   ├── Rubric ID
 │   ├── Criteria
 │   ├── Performance Levels
 │   └── Learning Objectives
 ├── Feedback (Entity)
 │   ├── Feedback ID
 │   ├── Feedback Content
 │   ├── Strengths
 │   ├── Areas for Improvement
 │   └── Evidence Reference
 └── Result (Entity)
     ├── Result ID
     ├── Score
     ├── Achievement Level
     ├── Competency Status
     └── Evidence Reference
```

**Why This Aggregate Root Matters**:
- Assessment is the measurement of learning progress
- Ensures consistency in how assessment is designed and conducted
- Provides single source of truth for all assessment data
- Enables traceability from assessment results to learning objectives and graduate profile dimensions
- Supports formative and summative assessment practices

**Invariants**:
- All Evidence must reference a valid Learning Activity
- All Rubrics must reference valid Learning Objectives
- All Feedback must reference valid Evidence
- All Results must reference valid Evidence
- Assessment must align with learning objectives and graduate profile dimensions

## Reporting Aggregate

```
Student Report (Aggregate Root)
 ├── Achievement (Entity)
 │   ├── Achievement ID
 │   ├── Subject
 │   ├── Learning Outcomes
 │   ├── Progress Level
 │   └── Assessment References
 ├── Narrative (Entity)
 │   ├── Narrative ID
 │   ├── Student Story
 │   ├── Strengths
 │   ├── Growth Areas
 │   ├── Dimension Progress
 │   └── Assessment References
 └── Recommendation (Entity)
     ├── Recommendation ID
     ├── Next Steps
     ├── Support Needs
     ├── Enrichment Opportunities
     └── Assessment References
```

**Why This Aggregate Root Matters**:
- Reporting is the communication of student progress
- Ensures consistency in how reports are generated and shared
- Provides single source of truth for all reporting data
- Enables traceability from reports to assessment results and graduate profile progress
- Supports stakeholder communication and decision-making

**Invariants**:
- All Achievement must reference valid Assessment data
- All Narratives must reference valid Assessment data
- All Recommendations must reference valid Assessment data
- Reports must align with graduate profile dimensions and learning objectives

## Additional Aggregate Roots

### Teacher Aggregate

```
Teacher (Aggregate Root)
 ├── Profile (Entity)
 ├── Certification (Entity)
 ├── Professional Development (Entity)
 └── Performance (Entity)
```

### Student Aggregate

```
Student (Aggregate Root)
 ├── Profile (Entity)
 ├── Enrollment (Entity)
 ├── Learning History (Entity)
 └── Portfolio (Entity)
```

### School Aggregate

```
School (Aggregate Root)
 ├── Profile (Entity)
 ├── Quality Data (Entity)
 ├── Improvement Plan (Entity)
 └── Accreditation (Entity)
```

## Aggregate Root Principles

### Consistency Boundaries
- Each Aggregate Root maintains consistency within its boundary
- No external entity can modify internal state without going through the Aggregate Root
- Aggregate Roots ensure business rules are enforced

### Transactional Boundaries
- Each Aggregate Root is a transactional boundary
- All changes within an aggregate happen in a single transaction
- Cross-aggregate operations require eventual consistency

### Identity and Lifecycle
- Each Aggregate Root has a unique identity
- Aggregate Roots have a lifecycle (created, modified, archived)
- Aggregate Root lifecycle is managed by the domain

### Access Patterns
- External entities access aggregates only through the Aggregate Root
- Internal entities are not directly accessible from outside the aggregate
- This encapsulation ensures consistency and enforces invariants

## Strategic Implications

### For Database Architecture
- Each Aggregate Root maps to a database aggregate or document
- Aggregate boundaries define transaction boundaries
- Indexing strategy follows aggregate access patterns

### For API Design
- Each Aggregate Root has its own API endpoints
- API operations respect aggregate boundaries
- Cross-aggregate operations require orchestration

### For AI Architecture
- AI agents operate within aggregate boundaries
- AI recommendations respect aggregate invariants
- AI-generated content follows aggregate structure

### For Development
- Each aggregate can be developed by a dedicated team
- Aggregate boundaries define module boundaries
- Testing strategy focuses on aggregate behavior

---

# SECTION 15.3 — Domain Events Architecture

## Event-Driven Architecture for NUSA

The Domain Events Architecture defines how domains communicate through events. Every important state change in NUSA generates a domain event that other domains can consume. This enables loose coupling, real-time integration, and AI agent orchestration.

## Core Principle

**Everything Important In NUSA Happens Through Domain Events**

Domain events are the primary mechanism for:
- Cross-domain communication
- AI agent triggering
- Data synchronization
- Real-time notifications
- Audit logging
- Analytics aggregation

## Domain Events Catalog

### Graduate Profile Events

#### GraduateProfileUpdated
- **Trigger**: Graduate Profile dimensions or indicators are updated
- **Producer Domain**: Graduate Profile Domain
- **Consumer Domain**: Curriculum Domain, Learning Domain, Assessment Domain, Reporting Domain
- **AI Agent Involved**: Curriculum AI Agent (updates curriculum alignment)

#### DimensionDefinitionApproved
- **Trigger**: New dimension definition is approved by national authorities
- **Producer Domain**: Graduate Profile Domain
- **Consumer Domain**: All domains
- **AI Agent Involved**: AI Orchestration (updates all domain references)

### Curriculum Events

#### CPPublished
- **Trigger**: New CP is published for a phase and subject
- **Producer Domain**: Curriculum Domain
- **Consumer Domain**: Learning Planning Domain, Assessment Domain
- **AI Agent Involved**: Curriculum AI Agent (generates TP recommendations)

#### TPGenerated
- **Trigger**: TP is generated from CP
- **Producer Domain**: Learning Planning Domain
- **Consumer Domain**: Learning Delivery Domain, Assessment Domain
- **AI Agent Involved**: Lesson Design Agent (generates ATP and Modul Ajar)

#### ATPApproved
- **Trigger**: ATP sequence is approved by teacher
- **Producer Domain**: Learning Planning Domain
- **Consumer Domain**: Learning Delivery Domain
- **AI Agent Involved**: Lesson Design Agent (generates Modul Ajar)

#### ModuleGenerated
- **Trigger**: Modul Ajar is generated from ATP
- **Producer Domain**: Learning Planning Domain
- **Consumer Domain**: Learning Delivery Domain
- **AI Agent Involved**: Lesson Design Agent (finalizes lesson plan)

### Learning Events

#### LearningActivityCompleted
- **Trigger**: Student completes a learning activity
- **Producer Domain**: Learning Delivery Domain
- **Consumer Domain**: Assessment Domain, Reporting Domain, Teacher Growth Domain
- **AI Agent Involved**: Teacher Copilot (provides reflection support)

#### StudentEngagementDetected
- **Trigger**: Student engagement level is detected during learning
- **Producer Domain**: Learning Delivery Domain
- **Consumer Domain**: Student Wellbeing Domain, Parent Partnership Domain
- **AI Agent Involved**: Teacher Copilot (suggests engagement strategies)

#### DeepLearningCycleCompleted
- **Trigger**: Student completes Understand-Apply-Reflect cycle
- **Producer Domain**: Learning Delivery Domain
- **Consumer Domain**: Assessment Domain, Reporting Domain
- **AI Agent Involved**: Pedagogy Coach (analyzes learning depth)

### Assessment Events

#### AssessmentSubmitted
- **Trigger**: Student submits assessment work
- **Producer Domain**: Assessment Domain
- **Consumer Domain**: Reporting Domain, Teacher Growth Domain
- **AI Agent Involved**: Assessment Agent (automated scoring)

#### AssessmentEvaluated
- **Trigger**: Assessment is evaluated by AI or teacher
- **Producer Domain**: Assessment Domain
- **Consumer Domain**: Reporting Domain, Learning Planning Domain
- **AI Agent Involved**: Assessment Agent (generates feedback)

#### FeedbackGenerated
- **Trigger**: Feedback is generated for student work
- **Producer Domain**: Assessment Domain
- **Consumer Domain**: Learning Delivery Domain, Reporting Domain, Parent Partnership Domain
- **AI Agent Involved**: Assessment Agent (personalized feedback)

#### LearningGapDetected
- **Trigger**: Learning gap is identified through assessment
- **Producer Domain**: Assessment Domain
- **Consumer Domain**: Learning Planning Domain, Teacher Growth Domain, Parent Partnership Domain
- **AI Agent Involved**: Assessment Agent (recommends remedial activities)

#### InterventionTriggered
- **Trigger**: Intervention is triggered based on assessment results
- **Producer Domain**: Assessment Domain
- **Consumer Domain**: Learning Planning Domain, Student Wellbeing Domain
- **AI Agent Involved**: Assessment Agent (generates intervention plan)

#### CompetencyMastered
- **Trigger**: Student demonstrates mastery of a competency
- **Producer Domain**: Assessment Domain
- **Consumer Domain**: Reporting Domain, Curriculum Domain
- **AI Agent Involved**: Assessment Agent (updates competency tracking)

### Reporting Events

#### ReportGenerated
- **Trigger**: Student report is generated
- **Producer Domain**: Reporting Domain
- **Consumer Domain**: Parent Partnership Domain, School Improvement Domain
- **AI Agent Involved**: Reporting Agent (narrative generation)

#### ParentSummaryCreated
- **Trigger**: Parent summary is created from report
- **Producer Domain**: Reporting Domain
- **Consumer Domain**: Parent Partnership Domain
- **AI Agent Involved**: Reporting Agent (translation and personalization)

#### GraduateProgressUpdated
- **Trigger**: Graduate profile progress is updated
- **Producer Domain**: Reporting Domain
- **Consumer Domain**: School Improvement Domain, Teacher Growth Domain
- **AI Agent Involved**: Reporting Agent (dimension progress analysis)

### Teacher Events

#### TeacherPDRecommended
- **Trigger**: Professional development is recommended for teacher
- **Producer Domain**: Teacher Professional Growth Domain
- **Consumer Domain**: Learning Delivery Domain, School Improvement Domain
- **AI Agent Involved**: PD Coach (resource matching)

#### WorkloadAlertTriggered
- **Trigger**: Teacher workload exceeds threshold
- **Producer Domain**: Teacher Performance & Workload Domain
- **Consumer Domain**: School Improvement Domain
- **AI Agent Involved**: Workload Advisor (support recommendations)

#### PerformanceEvaluationCompleted
- **Trigger**: Teacher performance evaluation is completed
- **Producer Domain**: Teacher Performance & Workload Domain
- **Consumer Domain**: Teacher Professional Growth Domain, School Improvement Domain
- **AI Agent Involved**: PD Coach (development planning)

### School Events

#### SchoolQualityMetricUpdated
- **Trigger**: School quality metric is updated
- **Producer Domain**: School Improvement Domain
- **Consumer Domain**: Quality Assurance & Accreditation Domain, Reporting Domain
- **AI Agent Involved**: School Intelligence Agent (analytics)

#### ImprovementStrategyApproved
- **Trigger**: School improvement strategy is approved
- **Producer Domain**: School Improvement Domain
- **Consumer Domain**: Teacher Professional Growth Domain, Curriculum Domain
- **AI Agent Involved**: School Intelligence Agent (implementation tracking)

#### AccreditationStatusChanged
- **Trigger**: School accreditation status changes
- **Producer Domain**: Quality Assurance & Accreditation Domain
- **Consumer Domain**: School Improvement Domain, Reporting Domain
- **AI Agent Involved**: Compliance Advisor (gap analysis)

### Parent Events

#### ParentEngagementRecorded
- **Trigger**: Parent engagement activity is recorded
- **Producer Domain**: Parent Partnership Domain
- **Consumer Domain**: Student Wellbeing Domain, Reporting Domain
- **AI Agent Involved**: Parent AI Agent (engagement analytics)

#### HomeLearningSupportRequested
- **Trigger**: Parent requests home learning support
- **Producer Domain**: Parent Partnership Domain
- **Consumer Domain**: Learning Planning Domain, Learning Delivery Domain
- **AI Agent Involved**: Parent AI Agent (resource recommendation)

### Student Events

#### WellbeingAlertTriggered
- **Trigger**: Student wellbeing alert is triggered
- **Producer Domain**: Student Wellbeing Domain
- **Consumer Domain**: Parent Partnership Domain, Teacher Growth Domain
- **AI Agent Involved**: Wellbeing Monitor (intervention recommendation)

#### CareerPathwaySelected
- **Trigger**: Student selects career pathway
- **Producer Domain**: Career & Future Readiness Domain
- **Consumer Domain**: Curriculum Domain, Learning Planning Domain
- **AI Agent Involved**: Career Counselor (pathway optimization)

### Platform Events

#### DataSyncCompleted
- **Trigger**: Data synchronization is completed across domains
- **Producer Domain**: Education Data & Interoperability Domain
- **Consumer Domain**: All domains
- **AI Agent Involved**: Data Steward (quality validation)

#### AIAgentExecuted
- **Trigger**: AI agent completes execution
- **Producer Domain**: AI Orchestration Domain
- **Consumer Domain**: All domains
- **AI Agent Involved**: AI Coordinator (orchestration logging)

#### AIRecommendationGenerated
- **Trigger**: AI recommendation is generated
- **Producer Domain**: AI Orchestration Domain
- **Consumer Domain**: Target domain based on recommendation type
- **AI Agent Involved**: Specific AI agent (recommendation delivery)

## Event Flow Patterns

### Upstream to Downstream Flow
Events flow from upstream contexts to downstream contexts:
- Graduate Profile Events → Curriculum Events → Learning Events → Assessment Events → Reporting Events

### Horizontal Flow
Events flow across horizontal contexts:
- Teacher Events ↔ School Improvement Events
- Parent Events ↔ Student Wellbeing Events
- Assessment Events ↔ Teacher Growth Events

### Cross-Domain Flow
Events flow across all domains through platform contexts:
- All domains → Data Events
- All domains → AI Orchestration Events
- AI Orchestration Events → All domains

## Event Processing Architecture

### Event Production
- Domain logic generates events when state changes
- Events are immutable and contain all relevant data
- Events are published to event bus

### Event Consumption
- Domains subscribe to relevant events
- Event handlers process events and trigger domain logic
- AI agents are triggered by specific events

### Event Storage
- All events are stored in event store for audit trail
- Event replay enables system recovery and analytics
- Event sourcing supports temporal queries

### AI Agent Integration
- AI agents subscribe to domain events
- AI agents process events and generate recommendations
- AI agents publish events for their actions

## Strategic Implications

### For System Architecture
- Event-driven architecture enables loose coupling
- Event bus is central integration mechanism
- Event sourcing provides audit trail and replay capability

### For AI Architecture
- AI agents are event-driven
- AI recommendations are triggered by domain events
- AI actions generate domain events

### For Data Architecture
- Event store provides complete audit trail
- Event replay enables data reconstruction
- Event streaming supports real-time analytics

### For Development
- Event-driven development enables independent deployment
- Event contracts define integration boundaries
- Event testing focuses on event flow and processing

---

# SECTION 15.4 — AI Agent ↔ Domain Mapping

## Complete AI Agent to Domain Mapping

This section provides the complete mapping between domains and AI agents, including human validators. This mapping ensures that every domain has appropriate AI support and human governance.

## Core Principles

**AI Generates**
**Human Governs**

**90% Operational Work**
**10% Human Validation**

## Domain-to-Agent Mapping Table

| Domain | AI Agent | Human Validator | Automation % | Validation Point |
|--------|----------|-----------------|--------------|------------------|
| **Graduate Profile** | Curriculum AI Agent | National Education Standards Experts | 80% | Dimension definitions and frameworks |
| **Foundational Competencies** | Curriculum AI Agent | Teachers and Parents | 70% | Competency assessments |
| **Curriculum** | Curriculum AI Agent | Teachers, Academic Coordinators | 90% | All curriculum materials |
| **Learning Planning** | Lesson Design Agent | Teachers | 90% | All lesson plans |
| **Learning Delivery** | Teacher Copilot | Teachers | 70% | All pedagogical decisions |
| **Deep Learning** | Pedagogy Coach | Teachers | 80% | All pedagogical decisions |
| **Assessment** | Assessment Agent | Teachers | 80% | All assessment items and high-stakes grading |
| **Reporting** | Reporting Agent | Homeroom Teachers | 95% | All reports |
| **Literacy** | Literacy Coach | Teachers | 80% | All literacy assessments |
| **Numeracy** | Math Coach | Teachers | 80% | All numeracy assessments |
| **Coding** | Coding Tutor | Teachers | 85% | All coding projects |
| **AI Literacy** | AI Guide | Teachers | 85% | All AI literacy projects |
| **Character Development** | Character Coach | Teachers and Parents | 70% | All character assessments |
| **Student Wellbeing** | Wellbeing Monitor | School Counselors and Parents | 75% | All wellbeing interventions |
| **Parent Partnership** | Parent AI Agent | Teachers | 85% | All parent communications |
| **Career & Future Readiness** | Career Counselor | School Counselors | 80% | All career guidance |
| **Lifelong Learning Record** | Learning Record Agent | Qualified Professionals | 85% | All credential verification and career recommendations |
| **Teacher Professional Growth** | PD Coach | Teachers and School Leaders | 85% | All PD recommendations |
| **Teacher Performance & Workload** | Workload Advisor | School Leaders | 80% | All performance evaluations |
| **School Improvement** | School Intelligence Agent | Principals | 85% | All improvement strategies |
| **Quality Assurance & Accreditation** | Compliance Advisor | Accrediting Bodies | 80% | All accreditation decisions |
| **Education Data & Interoperability** | Data Steward | Data Governance Board | 90% | All data governance decisions |
| **AI Orchestration** | AI Coordinator | AI Governance Board | 95% | All AI governance decisions |

## Detailed Domain Mappings

### Curriculum Domain

**AI Agent**: Curriculum AI Agent

**AI Responsibilities**:
- Generate TP from CP
- Generate ATP sequences
- Generate Modul Ajar
- Analyze curriculum alignment
- Recommend resources
- Perform gap analysis
- Generate personalized learning pathways

**Human Validator**: Teacher, Academic Coordinator

**Human Responsibilities**:
- Approve curriculum alignment
- Validate pedagogical quality
- Adapt to local context
- Select resources
- Ensure cultural relevance

**Human Validation Point**: All curriculum materials must be reviewed and approved by teachers

**Automation Target**: 90%

### Learning Planning Domain

**AI Agent**: Lesson Design Agent

**AI Responsibilities**:
- Generate TP from CP
- Generate ATP sequences
- Generate Modul Ajar
- Suggest differentiation strategies
- Match resources to learning objectives
- Optimize time allocation
- Support collaborative planning

**Human Validator**: Teacher

**Human Responsibilities**:
- Creative design
- Pedagogical decisions
- Relationship building
- Contextual adaptation
- Quality validation

**Human Validation Point**: All lesson plans must be reviewed and approved by teachers

**Automation Target**: 90%

### Assessment Domain

**AI Agent**: Assessment Agent

**AI Responsibilities**:
- Generate rubrics
- Generate feedback
- Generate remedial activities
- Automated scoring
- Pattern recognition
- Adaptive assessment
- Authentic evaluation
- Predictive analytics
- Multimodal assessment

**Human Validator**: Teacher

**Human Responsibilities**:
- Validate fairness
- Contextualize results
- Make final competency decisions
- High-stakes grading
- Ethical considerations

**Human Validation Point**: All assessment items must be validated for fairness; all high-stakes grading requires human approval

**Automation Target**: 80%

### Reporting Domain

**AI Agent**: Reporting Agent

**AI Responsibilities**:
- Generate narrative reports
- Generate parent summaries
- Data synthesis
- Progress communication
- Translation
- Visualization
- Anomaly detection
- Recommendations

**Human Validator**: Homeroom Teacher

**Human Responsibilities**:
- Validate student growth story
- Contextualize findings
- Adjust tone
- Ensure appropriateness
- Sensitive communication

**Human Validation Point**: All reports must be reviewed for accuracy and appropriateness before distribution

**Automation Target**: 95%

### School Improvement Domain

**AI Agent**: School Intelligence Agent

**AI Responsibilities**:
- School analytics
- School improvement planning
- Quality analytics
- Benchmarking
- Predictive modeling
- Recommendation engine
- Progress tracking
- Resource optimization
- Best practice matching

**Human Validator**: Principal

**Human Responsibilities**:
- Make decisions about improvement strategies
- Interpret analytics
- Lead change management
- Validate recommendations
- Community building

**Human Validation Point**: All improvement strategies must be validated by school leaders

**Automation Target**: 85%

### Teacher Professional Growth Domain

**AI Agent**: PD Coach

**AI Responsibilities**:
- Personalized learning recommendations
- Classroom analysis
- Resource matching
- Practice simulation
- Performance analytics
- Collaboration support
- Knowledge management
- PD recommendation

**Human Validator**: Teacher, School Leader

**Human Responsibilities**:
- Mentoring
- Support
- Relationship building
- Professional judgment
- Contextual interpretation

**Human Validation Point**: All PD recommendations must be validated by teachers and school leaders

**Automation Target**: 85%

### Parent Partnership Domain

**AI Agent**: Parent AI Agent

**AI Responsibilities**:
- Communication automation
- Translation
- Engagement analytics
- Resource recommendation
- Meeting scheduling
- Q&A
- Personalization
- Home learning support

**Human Validator**: Teacher

**Human Responsibilities**:
- Relationships
- Emotional support
- Cultural bridge
- Community building
- Decision-making

**Human Validation Point**: All parent communications must be reviewed by teachers

**Automation Target**: 85%

### Student Wellbeing Domain

**AI Agent**: Wellbeing Monitor

**AI Responsibilities**:
- Wellbeing monitoring
- Early warning
- Pattern recognition
- Intervention recommendations
- Progress tracking

**Human Validator**: School Counselor, Parent

**Human Responsibilities**:
- Emotional support
- Intervention decisions
- Contextual interpretation
- Ethical considerations
- Relationship building

**Human Validation Point**: All wellbeing interventions must be validated by school counselors and parents

**Automation Target**: 75%

### Career & Future Readiness Domain

**AI Agent**: Career Counselor

**AI Responsibilities**:
- Career matching
- Pathway recommendation
- Skill gap analysis
- Industry trend analysis
- Personalized guidance

**Human Validator**: School Counselor

**Human Responsibilities**:
- Mentoring
- Relationship building
- Contextual interpretation
- Ethical guidance
- Decision support

**Human Validation Point**: All career guidance must be validated by school counselors

**Automation Target**: 80%

### Lifelong Learning Record Domain

**AI Agent**: Learning Record Agent

**AI Responsibilities**:
- Competency gap analysis
- Career recommendation
- Personalized learning path
- Workforce readiness prediction
- Learning pattern analysis
- Credential verification
- Portfolio curation

**Human Validator**: Qualified Professionals

**Human Responsibilities**:
- Record validation
- Credential verification
- Privacy decisions
- Career guidance
- Contextual interpretation
- Ethical considerations

**Human Validation Point**: All credential verification and career recommendations must be validated by qualified professionals

**Automation Target**: 85%

## AI Agent Coordination

### AI Orchestration Module

The AI Orchestration Domain coordinates all AI agents across domains:

- **AI Coordinator**: Manages agent lifecycle and coordination
- **Event-Driven Triggers**: AI agents are triggered by domain events
- **Cross-Domain Collaboration**: AI agents collaborate across domain boundaries
- **Human-in-the-Loop**: All AI agent outputs require human validation at defined points

### Agent Interaction Patterns

1. **Sequential**: One agent's output becomes another agent's input (e.g., Curriculum AI Agent → Lesson Design Agent)
2. **Parallel**: Multiple agents work on different aspects simultaneously (e.g., Literacy Coach and Math Coach)
3. **Hierarchical**: Higher-level agents coordinate lower-level agents (e.g., AI Coordinator → Domain Agents)
4. **Collaborative**: Agents share information and coordinate actions (e.g., Teacher Copilot and Assessment Agent)

## Human Governance Framework

### Validation Points

Human validation occurs at critical points:
- **Before Action**: AI-generated content is validated before being used
- **After Action**: AI-recommended actions are reviewed after execution
- **Periodic**: AI agent performance is reviewed periodically
- **Exception**: Human intervention is triggered for exceptions or anomalies

### Validation Roles

- **Teacher**: Validates curriculum, lesson plans, assessments, reports, parent communications
- **Academic Coordinator**: Validates curriculum alignment and quality
- **School Leader**: Validates improvement strategies, performance evaluations, PD recommendations
- **School Counselor**: Validates wellbeing interventions, career guidance
- **Parent**: Validates character assessments, wellbeing interventions
- **National Expert**: Validates graduate profile dimensions and frameworks
- **Data Governance Board**: Validates data governance decisions
- **AI Governance Board**: Validates AI governance decisions
- **Accrediting Body**: Validates accreditation decisions

## Strategic Implications

### For AI Architecture
- Each domain has dedicated AI agent(s)
- AI agents are event-driven and coordinated centrally
- Human validation is built into AI agent workflows
- AI agent performance is continuously monitored

### For User Experience
- AI agents provide context-aware support
- AI recommendations are transparent and explainable
- Human validation is seamless and non-intrusive
- Users maintain control and oversight

### For Governance
- Human validation ensures accountability
- AI governance board oversees AI agent behavior
- Data governance board ensures data integrity
- Regular audits ensure compliance and quality

---

# SECTION 15.5 — Competency Graph Architecture

## National Competency Graph

NUSA does not merely store grades. NUSA stores the national competency graph that represents the complete knowledge and skill structure of Indonesian education.

## Competency Node

Competency nodes represent specific skills, knowledge, or abilities that students develop throughout their educational journey.

**Examples of Competency Nodes**:
- Critical Thinking
- Collaboration
- Creativity
- Numeracy
- Coding
- AI Literacy
- Scientific Reasoning
- Digital Citizenship
- Problem Solving
- Communication

**Node Structure**:
- Competency ID
- Competency Name
- Competency Description
- Developmental Phases
- Assessment Criteria
- Prerequisite Competencies
- Related Competencies

## Competency Relationship

Competencies are not isolated. They form a graph structure with relationships that define how competencies build upon each other.

**Relationship Types**:

### Prerequisite Relationship
A competency that must be mastered before another competency can be developed.

```
Numeracy
→ prerequisite →
Data Analysis

Data Analysis
→ prerequisite →
AI Literacy
```

### Related Relationship
Competencies that are related and reinforce each other.

```
Critical Thinking
↔ related →
Problem Solving
```

### Part-Of Relationship
A competency that is part of a broader competency.

```
Addition
→ part-of →
Arithmetic Operations
```

### Enables Relationship
A competency that enables the development of another competency.

```
Reading Comprehension
→ enables →
Research Skills
```

## Mastery State

Each competency node has a mastery state that represents the student's current level of development.

**Mastery States**:
- **Not Started**: Student has not begun developing this competency
- **Emerging**: Student is beginning to show initial understanding
- **Developing**: Student is actively developing the competency with support
- **Proficient**: Student demonstrates consistent competency with minimal support
- **Advanced**: Student demonstrates mastery and can apply the competency independently in complex contexts

**Mastery Tracking**:
- Current mastery level
- Mastery trajectory over time
- Evidence of mastery
- Areas for development
- Recommended next steps

## Competency Graph Purpose

The competency graph serves multiple critical purposes in NUSA:

### Personalization
- Identify student's current position in the competency graph
- Recommend personalized learning pathways
- Adapt learning activities to student's competency level
- Provide targeted support for specific competencies

### Recommendation
- Suggest next competencies to develop based on mastery state
- Recommend resources matched to competency level
- Provide learning pathway optimization
- Suggest peer collaboration based on complementary competencies

### Intervention
- Identify students at risk of competency regression
- Trigger early intervention when mastery declines
- Provide targeted remedial activities
- Monitor intervention effectiveness

### Pathway Generation
- Generate personalized learning pathways
- Optimize pathway based on student's goals and constraints
- Provide alternative pathways when needed
- Support lifelong learning pathway planning

### AI Tutoring
- AI tutors use competency graph to understand student's knowledge state
- Provide targeted explanations based on competency gaps
- Adapt tutoring strategies to mastery level
- Connect new learning to existing competency structure

## Competency Graph Architecture

### Graph Structure
- **Nodes**: Individual competencies
- **Edges**: Relationships between competencies
- **Weights**: Strength of relationships
- **Layers**: Competency levels (foundational, intermediate, advanced)

### Graph Properties
- **Directed**: Relationships have direction (prerequisite, enables, etc.)
- **Weighted**: Relationships have strength indicators
- **Multi-layered**: Competencies exist at different levels of abstraction
- **Dynamic**: Graph evolves as new competencies are identified

### Graph Operations
- **Traversal**: Find pathways between competencies
- **Query**: Find competencies meeting specific criteria
- **Update**: Add or modify competencies and relationships
- **Analysis**: Identify competency clusters and patterns

## Strategic Implications

### For AI Architecture
- Competency graph is a core intelligence layer
- AI agents use graph for reasoning and recommendation
- Graph enables sophisticated personalization at scale
- Graph supports AI tutoring and adaptive learning

### For Data Architecture
- Graph database for efficient graph operations
- Competency data as first-class citizen
- Mastery state tracking with temporal history
- Graph analytics for insights and patterns

### For Curriculum Design
- Curriculum mapped to competency graph
- Learning objectives aligned with competency nodes
- Curriculum sequences follow prerequisite relationships
- Curriculum gaps identified through graph analysis

### For Assessment
- Assessments mapped to competency nodes
- Mastery state updated through assessment results
- Assessment recommendations based on competency gaps
- Competency growth tracked over time

## Core Principle

**Competency Graph is one of the core intelligence layers of NUSA.**

The competency graph enables NUSA to move beyond simple grade tracking to sophisticated understanding of student knowledge and skill development. This is foundational for AI-native education at scale.

---

# SECTION 15.6 — Learning Digital Twin Architecture

## Learning Digital Twin Concept

Learning Digital Twin is the digital representation of each student's learning journey and development. It is not an avatar, not a chatbot, and not an administrative profile. It is a comprehensive digital representation that captures the complete learning state, history, and trajectory of each student.

## Digital Twin Components

### Identity
- Student identification
- Demographic information
- Educational context
- Learning preferences
- Special needs and accommodations

### Learning History
- All learning activities completed
- Time spent on activities
- Engagement patterns
- Learning pathways taken
- Historical performance

### Competency State
- Current mastery levels for all competencies
- Competency growth trajectory
- Competency gaps
- Strengths and areas for development
- Prerequisite completion status

### Assessment Evidence
- All assessment results
- Evidence of learning
- Performance trends
- Assessment patterns
- Growth indicators

### Portfolio Evidence
- Artifacts and projects
- Creative work
- Research outputs
- Presentations and performances
- Peer collaboration evidence

### Learning Preferences
- Preferred learning modalities
- Content preferences
- Interaction preferences
- Time preferences
- Environmental preferences

### Intervention History
- Interventions received
- Intervention effectiveness
- Support services accessed
- Response to interventions
- Ongoing support needs

### Growth Trajectory
- Predicted learning outcomes
- Recommended pathways
- Career readiness indicators
- Graduate profile progress
- Future learning opportunities

## Purpose

The Learning Digital Twin serves multiple critical purposes:

### Personalization Foundation
- Enables truly personalized learning at scale
- Supports adaptive learning systems
- Provides context for AI recommendations
- Enables individualized pathway generation

### Continuous Understanding
- Maintains comprehensive understanding of each student
- Tracks development over time
- Identifies patterns and trends
- Supports longitudinal analysis

### Intervention Support
- Identifies students needing support early
- Provides context for intervention design
- Monitors intervention effectiveness
- Supports targeted support

### Parent Communication
- Provides comprehensive view for parents
- Enables meaningful parent engagement
- Supports home-school collaboration
- Facilitates informed decision-making

## Inputs

The Learning Digital Twin is continuously updated with data from:

- Learning activities and engagement
- Assessment results and evidence
- Portfolio artifacts and reflections
- Teacher observations and feedback
- Peer interactions and collaboration
- Parent feedback and observations
- Wellbeing and engagement data
- External learning experiences

## Outputs

The Learning Digital Twin generates:

- Personalized learning recommendations
- Adaptive learning pathways
- Intervention alerts and recommendations
- Progress reports and insights
- Parent communication summaries
- Teacher support recommendations
- Career readiness assessments
- Graduate profile progress tracking

## AI Opportunities

The Learning Digital Twin enables sophisticated AI capabilities:

### Predictive Analytics
- Predict learning outcomes
- Identify at-risk students early
- Recommend proactive interventions
- Optimize learning pathways

### Adaptive Learning
- Dynamically adapt content difficulty
- Personalize learning pace
- Adjust learning modalities
- Provide targeted support

### Intelligent Tutoring
- AI tutors understand complete learning context
- Provide personalized explanations
- Connect new learning to existing knowledge
- Adapt tutoring strategies to individual needs

### Recommendation Engine
- Recommend learning resources
- Suggest learning activities
- Optimize learning sequences
- Provide peer collaboration opportunities

## Relationships

### Assessment Domain
- Digital Twin consumes assessment evidence
- Assessment updates competency state
- Digital Twin provides context for assessment design
- Assessment results update growth trajectory

### Portfolio Domain
- Digital Twin incorporates portfolio evidence
- Portfolio artifacts demonstrate competency mastery
- Digital Twin provides portfolio curation guidance
- Portfolio supports competency verification

### Recommendation Engine
- Digital Twin provides context for recommendations
- Recommendations update based on digital twin state
- Digital Twin learns from recommendation effectiveness
- Recommendations influence learning trajectory

### AI Tutor
- AI Tutor uses digital twin for personalization
- Tutor interactions update digital twin
- Digital twin provides tutoring context
- Tutor effectiveness tracked in digital twin

### Parent Partnership
- Digital Twin provides parent communication content
- Parent feedback updates digital twin
- Digital twin supports parent engagement
- Parent insights inform learning trajectory

## Strategic Implications

### For Personalization at Scale
- Digital twin enables individualized learning for millions of students
- AI can reason about complete learning context
- Personalization is data-driven and evidence-based
- Scale does not compromise personalization quality

### For AI Architecture
- Digital twin is the context layer for AI agents
- AI agents use digital twin for reasoning
- Digital twin enables sophisticated AI capabilities
- AI effectiveness improves with digital twin maturity

### For Data Architecture
- Digital twin requires comprehensive data integration
- Data quality directly impacts digital twin accuracy
- Real-time data updates enable real-time personalization
- Data privacy and security are critical

### For Student Empowerment
- Students own their digital twin
- Digital twin supports student agency
- Students can track their own growth
- Digital twin enables informed learning decisions

## Core Principle

**Learning Digital Twin becomes the foundation for personalization at national scale.**

The digital twin enables NUSA to provide truly personalized education for millions of students while maintaining the human-centered approach that defines quality education. This is foundational for the AI-Native Education Platform paradigm.

---

# SECTION 15.7 — Education Capability Architecture

## Three-Layer Separation

In NUSA, there are three distinct layers that must not be mixed. Each layer serves a specific purpose and operates at different levels of abstraction and stability.

## Layer 1 — Domain

Domains represent the business areas of education. These are stable, long-lasting structures that define the fundamental boundaries of the education system.

**Examples**:
- Curriculum Domain
- Assessment Domain
- Teacher Growth Domain
- Parent Partnership Domain
- Student Wellbeing Domain
- School Improvement Domain
- Lifelong Learning Record Domain

**Characteristics**:
- Stable for decades
- Based on educational theory and policy
- Independent of technology implementation
- Define business boundaries and ownership
- Serve as the foundation for all other layers

**Purpose**: Domains own the business logic and define the fundamental structure of the education system.

## Layer 2 — Capability

Capabilities are operational capabilities that the system executes. These represent what the system can do and are more flexible than domains.

**Examples**:
- Curriculum Planning
- Competency Mapping
- Assessment Analysis
- Personalized Learning
- Learning Recommendation
- Teacher Coaching
- Parent Communication
- School Quality Analytics
- Career Pathway Planning

**Characteristics**:
- Can change based on organizational needs
- Implement domain business logic
- More flexible than domains
- Technology-agnostic
- Focused on operational execution

**Purpose**: Capabilities execute the work defined by domains.

## Layer 3 — AI Services

AI Services are AI capabilities that accelerate and enhance capabilities. These are not domains but a service layer that provides intelligent assistance.

**Examples**:
- Content Generation
- Recommendation Engine
- Predictive Analytics
- Natural Language Processing
- Learning Personalization
- Narrative Generation
- Pattern Recognition
- Adaptive Learning

**Characteristics**:
- Rapidly evolving
- Technology-dependent
- Service-oriented
- Accelerate capabilities
- Can be replaced or upgraded independently

**Purpose**: AI accelerates the capabilities.

## Core Principle

> Domains own the business logic.
>
> Capabilities execute the work.
>
> AI accelerates the capabilities.

This separation ensures that:
- Business logic remains stable even as AI technology evolves
- AI can be upgraded without changing domain structure
- Capabilities can be reconfigured without affecting domain boundaries
- The system remains flexible and adaptable while maintaining architectural integrity

## Strategic Implications

### For Architecture
- Clear separation of concerns enables independent evolution
- Domains provide stable foundation for all other layers
- Capabilities bridge domains and AI services
- AI services can be swapped without affecting business logic

### For Development
- Teams can work on different layers independently
- AI development does not require domain restructuring
- Capability changes do not require domain changes
- Technology choices are isolated to AI service layer

### For Governance
- Domain ownership is clear and stable
- Capability ownership can be flexible
- AI governance is focused on service layer
- Human authority is maintained at domain and capability levels

---

# SECTION 15.8 — Lifelong Learning Record Architecture

## National Learning Identity

NUSA does not stop at school. Every individual has a National Learning Identity that stores their complete learning journey throughout their entire life.

## Learning Lifecycle

The lifelong learning record spans all educational phases:

```
PAUD (Early Childhood Education)
    ↓
SD (Primary School)
    ↓
SMP (Junior High School)
    ↓
SMA/SMK (Senior High School/Vocational School)
    ↓
Higher Education (University/College)
    ↓
Professional Certification
    ↓
Workplace Learning
    ↓
Lifelong Learning (Continuous Upskilling/Reskilling)
```

## Record Components

The lifelong learning record stores comprehensive evidence of learning:

### Competencies
- All competencies developed across all phases
- Mastery levels and progression over time
- Competency gaps and areas for development
- Cross-phase competency development

### Achievements
- Academic achievements
- Extracurricular achievements
- Competition results
- Recognition and awards
- Milestone accomplishments

### Portfolios
- Project artifacts
- Creative work
- Research outputs
- Presentations and performances
- Collaborative work evidence

### Certifications
- Professional certifications
- Industry credentials
- Skill certifications
- Language proficiency certificates
- Specialized training certificates

### Projects
- Individual projects
- Team projects
- Community service projects
- Innovation projects
- Entrepreneurship initiatives

### Skills
- Technical skills
- Soft skills
- Digital skills
- Industry-specific skills
- Transferable skills

### Character Evidence
- Character development evidence
- Values demonstration
- Leadership experiences
- Community engagement
- Ethical behavior evidence

## Snapshot vs Longitudinal History

### Report Cards are Snapshots
- Capture performance at a specific point in time
- Limited to a single educational phase
- Often lost when students transition
- Do not provide comprehensive learning history
- Limited to academic metrics

### Learning Record is Longitudinal History
- Captures complete learning journey across all phases
- Maintains continuity throughout life
- Portable across institutions and phases
- Provides comprehensive evidence of development
- Includes academic, character, and skill development

## Strategic Importance

### For Individuals
- Own their complete learning history
- Can demonstrate comprehensive development
- Support career advancement and transitions
- Enable lifelong learning planning
- Provide evidence for opportunities

### For Employers
- Access comprehensive candidate profiles
- Verify skills and competencies
- Make informed hiring decisions
- Support workforce planning
- Enable skills-based hiring

### For Education System
- Track learning outcomes across phases
- Identify learning gaps and strengths
- Support curriculum improvement
- Enable evidence-based policy
- Measure system effectiveness

### For National Development
- Build national skills inventory
- Support workforce planning
- Enable human capital development
- Support Indonesia Emas 2045 goals
- Provide data for national education strategy

## Core Principle

**Rapor hanyalah snapshot. Learning Record adalah longitudinal history.**

The lifelong learning record transforms education from a series of disconnected snapshots into a continuous, comprehensive journey that supports individuals throughout their entire lives and provides the nation with the data needed for strategic human capital development.

---

# SECTION 15.9 — AI Orchestration Architecture

## Multi-Agent System

NUSA is not a single AI. NUSA uses multiple AI Agents, each specialized for specific educational tasks. These agents work together under orchestration to provide comprehensive AI-native education experience.

## AI Agent Examples

### Curriculum Agent

**Purpose**: Assist teachers in curriculum planning and development

**Capabilities**:
- CP → TP transformation
- TP → ATP generation
- Curriculum alignment verification
- Learning objective refinement
- Resource recommendation for curriculum

**Human Validation**: Teachers review and approve all curriculum outputs

---

### Assessment Agent

**Purpose**: Support assessment design, administration, and analysis

**Capabilities**:
- Rubric analysis and improvement
- Evidence classification and tagging
- Assessment item generation
- Performance pattern analysis
- Competency assessment support

**Human Validation**: Teachers review assessment design and interpret results

---

### Reporting Agent

**Purpose**: Automate report generation and communication

**Capabilities**:
- Narrative report generation
- Data synthesis and summarization
- Multi-language translation
- Report customization
- Communication formatting

**Human Validation**: Teachers review and approve all reports before sending

---

### Teacher Coach Agent

**Purpose**: Support teacher professional development

**Capabilities**:
- Professional growth recommendation
- Professional development resource matching
- Classroom practice analysis
- Teaching strategy suggestions
- Workload optimization recommendations

**Human Validation**: Teachers and school leaders review and act on recommendations

---

### Parent Engagement Agent

**Purpose**: Facilitate parent communication and engagement

**Capabilities**:
- Parent communication automation
- Translation and localization
- Engagement analytics
- Communication timing optimization
- Parent question answering

**Human Validation**: Teachers review and approve all parent communications

---

### Learning Coach Agent

**Purpose**: Support personalized learning for students

**Capabilities**:
- Learning pathway recommendation
- Personalized resource matching
- Learning gap identification
- Study strategy suggestions
- Motivation and engagement support

**Human Validation**: Teachers oversee learning pathways and intervene as needed

---

### Career Agent

**Purpose**: Support career and future readiness

**Capabilities**:
- Career pathway recommendation
- Skill gap analysis
- Industry trend analysis
- Opportunity matching
- Future skill prediction

**Human Validation**: School counselors review and guide career decisions

---

### Wellbeing Agent

**Purpose**: Monitor and support student wellbeing

**Capabilities**:
- Engagement pattern monitoring
- Early warning for wellbeing issues
- Intervention recommendation
- Support resource matching
- Wellbeing trend analysis

**Human Validation**: School counselors and parents review and act on wellbeing alerts

---

### School Improvement Agent

**Purpose**: Support school quality improvement

**Capabilities**:
- School quality analytics
- Benchmarking analysis
- Improvement strategy recommendation
- Resource optimization suggestions
- Quality trend monitoring

**Human Validation**: Principals and school leaders review and implement improvement strategies

---

### Data Intelligence Agent

**Purpose**: Provide analytics and insights

**Capabilities**:
- Learning analytics
- Predictive modeling
- Pattern recognition
- Trend analysis
- Insight generation

**Human Validation**: Data governance board reviews analytics and insights

---

### AI Orchestrator Agent

**Purpose**: Coordinate all AI agents

**Capabilities**:
- Agent lifecycle management
- Cross-agent coordination
- Workflow orchestration
- Event-driven agent triggering
- Agent performance monitoring

**Human Validation**: AI governance board oversees agent behavior and coordination

---

## Core Principle

```text
AI Agents do not own decisions.
AI Agents provide recommendations.
Humans own accountability.
```

This principle ensures that:
- AI accelerates but does not replace human judgment
- Educational accountability remains with humans
- AI provides data-driven recommendations
- Humans make final educational decisions
- AI supports but does not govern

## Strategic Implications

### For AI Architecture
- Multi-agent system enables specialization
- Agents can be developed and updated independently
- Orchestrator ensures coordinated agent behavior
- Human validation points maintain accountability

### For Education Quality
- AI provides data-driven support
- Humans maintain educational judgment
- Recommendations are based on comprehensive data
- Human expertise guides AI learning

### For Governance
- Clear human authority at decision points
- AI governance board oversees agent behavior
- Regular audits ensure AI alignment
- Human-in-the-loop maintains accountability

---

# SECTION 15.10 — Human Governance Layer

## 90% AI / 10% Human Target

The target of 90% AI assistance and 10% human judgment does not mean eliminating humans. It means that AI handles routine, data-intensive, and repetitive tasks, while humans focus on high-value educational decisions that require judgment, ethics, and emotional intelligence.

## Decision Levels

### Level 1: AI Autonomous

Decisions that AI can make independently without human intervention.

**Examples**:
- Scheduling and calendar management
- Document formatting and layout
- Routine document generation
- Data aggregation and summarization
- Pattern recognition and classification
- Translation and localization

**Rationale**: These tasks are well-defined, rule-based, and have clear success criteria. AI can perform them efficiently and accurately.

**Human Role**: Monitor for quality and intervene if issues arise.

---

### Level 2: AI Assisted

Decisions where AI provides recommendations that humans review and approve.

**Examples**:
- ATP (Learning Path) recommendation
- Assessment item recommendation
- Intervention strategy recommendation
- Professional development recommendation
- Resource recommendation
- Learning pathway recommendation

**Rationale**: These tasks require educational judgment and context. AI provides data-driven recommendations, but humans make final decisions.

**Human Role**: Review recommendations, apply educational judgment, approve or modify.

---

### Level 3: Human Approval Required

Decisions where AI provides analysis but humans must explicitly approve.

**Examples**:
- Final grades
- Graduation decisions
- Student interventions
- Disciplinary actions
- Promotion decisions
- Certification awards

**Rationale**: These decisions have significant impact on students' lives and futures. Human approval ensures accountability and fairness.

**Human Role**: Review AI analysis, apply ethical and educational judgment, make final decision with full accountability.

---

### Level 4: Human Only

Decisions that only humans can make.

**Examples**:
- Ethical judgment
- Emotional mentoring
- Conflict resolution
- Complex disciplinary decisions
- Sensitive family situations
- Crisis management

**Rationale**: These decisions require emotional intelligence, empathy, ethical reasoning, and complex judgment that AI cannot provide.

**Human Role**: Complete human decision-making with no AI involvement in the decision itself.

---

## Core Principle

> Educational accountability can never be delegated entirely to AI.

This principle ensures that:
- Humans maintain ultimate accountability for educational outcomes
- AI supports but does not replace human educational judgment
- Ethical and emotional dimensions remain human responsibility
- High-stakes decisions always require human approval
- The education system remains human-centered

## Strategic Implications

### For AI Development
- Focus AI on tasks where it provides clear value
- Design clear human validation points
- Ensure AI recommendations are explainable
- Maintain human oversight capabilities

### For Education Quality
- Humans apply educational judgment to AI recommendations
- AI provides data-driven insights
- Humans ensure ethical and equitable application
- Human expertise guides continuous improvement

### For Governance
- Clear definition of decision levels
- Human authority maintained at critical points
- Regular review of AI-human decision boundaries
- Accountability remains with humans

---

# SECTION 15.11 — Outcome Traceability Matrix

## Traceability Hierarchy

Every educational activity in NUSA must be traceable from the highest national vision to the most granular learning activity. This ensures that everything contributes to the ultimate goal of Indonesia Emas 2045.

## Outcome Traceability Table

| Level | Description | Purpose |
|-------|-------------|---------|
| **National Vision** | Indonesia Emas 2045 | Ultimate national goal for human capital development |
| **Graduate Profile** | 8 Dimensi Profil Lulusan | National graduate profile defining desired outcomes |
| **Learning Outcome** | Kompetensi per jenjang | Competencies defined for each educational phase |
| **CP** | Capaian Pembelajaran | Learning capabilities defined in curriculum |
| **TP** | Tujuan Pembelajaran | Learning objectives derived from CP |
| **ATP** | Learning Path | Sequence of learning activities |
| **Learning Activity** | Experience | Specific learning experiences for students |
| **Assessment** | Evidence Collection | Assessment activities to collect evidence |
| **Competency Achievement** | Student Progress | Demonstration of competency mastery |
| **Learning Record** | Lifelong Learning History | Complete longitudinal learning record |

## Traceability Principles

```text
Every activity must contribute to a competency.
Every competency must contribute to an outcome.
Every outcome must contribute to the Graduate Profile.
Every Graduate Profile dimension must contribute to Indonesia Emas 2045.
```

## Traceability Flow

```
Indonesia Emas 2045
    ↓ defines
Graduate Profile (8 Dimensions)
    ↓ drives
Learning Outcomes (per phase)
    ↓ informs
CP (Capaian Pembelajaran)
    ↓ breaks down into
TP (Tujuan Pembelajaran)
    ↓ sequences into
ATP (Learning Path)
    ↓ delivers
Learning Activities
    ↓ assessed through
Assessment
    ↓ demonstrates
Competency Achievement
    ↓ recorded in
Learning Record (Lifelong)
```

## Strategic Importance

### For Alignment
- Ensures all activities contribute to national goals
- Maintains focus on Graduate Profile development
- Prevents misalignment between activities and outcomes
- Enables strategic coherence across the system

### For Accountability
- Every activity can be traced to its purpose
- Resource allocation can be justified
- Impact can be measured at all levels
- Continuous improvement is data-driven

### For AI
- AI agents understand the complete context
- Recommendations are aligned with outcomes
- Personalization supports outcome achievement
- AI effectiveness can be measured against outcomes

### For Governance
- Clear line of sight from activities to goals
- Decision-making is outcome-focused
- Policy implementation is traceable
- System performance can be evaluated

## Core Principle

**Every educational decision must be traceable through the Competency Graph.**

This principle ensures that:
- No activity is disconnected from educational purpose
- Resources are allocated to outcome-achieving activities
- AI recommendations are aligned with educational goals
- The system maintains coherence and focus

---

# SECTION 16 — National Education Domain Coverage

## Purpose

This section demonstrates that the domain model comprehensively covers all national education functions, ensuring that the Education Operating System (NUSA) can serve as the complete foundation for Indonesian education.

## National Education Function Coverage Matrix

| National Education Function | Domain Owner | Description |
| --------------------------- | ------------ | ----------- |
| **National Standards** | Graduate Profile Domain | National graduate profile (8 Dimensions) defining desired outcomes |
| **Graduate Profile** | Graduate Profile Domain | Implementation and tracking of Profil Lulusan 8 Dimensi |
| **Curriculum** | Curriculum Domain | National curriculum (CP) as single source of truth |
| **Learning Planning** | Learning Planning Domain | TP, ATP, and Modul Ajar generation and management |
| **Learning Delivery** | Learning Domain | Learning activities, experiences, and delivery |
| **Assessment** | Assessment Domain | Formative and summative assessment, evidence collection |
| **Reporting** | Reporting Domain | Student progress reporting to parents and stakeholders |
| **Literacy** | Literacy Domain | Reading and writing development across all phases |
| **Numeracy** | Numeracy Domain | Mathematical thinking and problem-solving development |
| **Coding & AI** | Coding & AI Literacy Domain | Digital skills, coding, and AI literacy development |
| **Character Development** | Character Development Domain | Values, ethics, and character education |
| **Student Development** | Student Wellbeing Domain | Physical, mental, and emotional wellbeing support |
| **Special Education** | Learning Domain (inclusive) | Support for students with special needs |
| **Parent Partnership** | Parent Partnership Domain | Parent engagement and communication |
| **Teacher Growth** | Teacher Professional Growth Domain | Professional development and teacher capacity building |
| **School Improvement** | School Improvement Domain | Continuous school quality improvement |
| **School Leadership** | School Improvement Domain | Leadership development and school management |
| **School Quality Assurance** | Quality Assurance & Accreditation Domain | Accreditation and quality monitoring |
| **Career Readiness** | Career & Future Readiness Domain | Career guidance and future pathway planning |
| **Industry Partnership** | Career & Future Readiness Domain | Industry-education collaboration |
| **Lifelong Learning** | Lifelong Learning Record Domain | National learning identity and lifelong learning record |
| **Education Analytics** | Education Data & Interoperability Domain | Learning analytics and data intelligence |
| **AI Governance** | AI Orchestration Domain | AI agent governance and human-in-the-loop oversight |

## Domain Completeness Statement

This domain model provides comprehensive coverage of the Indonesian national education system:

### Complete Process Coverage
- All core educational processes are represented within the domain model
- From national standards to classroom implementation
- From student enrollment to lifelong learning
- From teacher preparation to professional development
- From school operations to national quality assurance

### No Strategic Gaps
- No strategic area of national education exists outside this domain model
- All critical education functions have designated domain ownership
- Cross-domain relationships are explicitly defined
- Integration points between domains are architecturally specified

### Foundation for National NUSA
- This domain model can serve as the complete foundation for the national Education Operating System
- Scalable to serve millions of students across all educational phases
- Designed to support the vision of Indonesia Emas 2045
- Aligned with Permendikdasmen 2025 and national education policy

### Architectural Integrity
- Domain boundaries are clear and stable
- Domain relationships are well-defined
- Domain ownership is unambiguous
- Domain evolution is governed and controlled

## Strategic Implications

### For National Implementation
- NUSA can be implemented nationally based on this domain model
- No additional domains are required for complete coverage
- Existing domains can be extended without architectural restructuring
- Domain model supports phased national rollout

### For System Architecture
- Complete domain coverage enables comprehensive system design
- All system components can be mapped to domains
- Data architecture can be designed based on domain boundaries
- AI agent architecture can be aligned with domain responsibilities

### For Governance
- Clear domain ownership enables effective governance
- Domain boundaries support accountability structures
- Cross-domain coordination is architecturally enabled
- Human authority is maintained at domain level

---

# SECTION 17 — National Learning Intelligence Backbone

## Strategic Intelligence Components

The Education Operating System (NUSA) is built upon three strategic intelligence components that form the backbone of the entire system. These components work together to enable AI-native education at national scale while maintaining human-centered educational values.

## Layer 1 — Competency Graph

### Definition

The Competency Graph is the comprehensive representation of all national competencies, their relationships, and the learning pathways that connect them.

### Components

- **Competency Nodes**: All competencies defined in the national education system
- **Competency Relationships**: Prerequisite, related, part-of, and enables relationships
- **Prerequisite Map**: Clear mapping of competency dependencies
- **Learning Progression Map**: Optimal pathways for competency development

### Example Flow

```
Graduate Profile (8 Dimensions)
    ↓ breaks down into
CP (Capaian Pembelajaran)
    ↓ defines
TP (Tujuan Pembelajaran)
    ↓ sequences into
Learning Activities
    ↓ assessed through
Assessment Evidence
    ↓ demonstrates
Competency Achievement
```

### Strategic Role

**The Competency Graph is "The Educational Brain of NUSA."**

It serves as:
- Single Source of Truth for all educational competencies
- Foundation for recommendation engines
- Basis for AI personalization
- Core of learning analytics
- Enabler of outcome traceability

---

## Layer 2 — Learning Digital Twin

### Definition

The Learning Digital Twin is the digital representation of each student's complete learning journey and development state.

### Components

- **Competency Mastery**: Current mastery levels for all competencies
- **Learning Behavior**: Patterns of engagement and learning preferences
- **Assessment History**: Complete record of all assessments and results
- **Portfolio Evidence**: Artifacts, projects, and creative work
- **Interests**: Student interests and passions
- **Strengths**: Areas where student demonstrates excellence
- **Learning Gaps**: Areas requiring additional support

### Architectural Relationship

**Learning Digital Twin is built upon the Competency Graph.**

The Competency Graph provides:
- The structure for organizing competency mastery data
- The relationships for understanding learning progression
- The pathways for personalized learning recommendations
- The context for interpreting assessment evidence

The Learning Digital Twin provides:
- The student-specific data for each competency node
- The behavioral patterns that inform personalization
- The evidence that demonstrates competency achievement
- The trajectory that predicts future learning needs

---

## Layer 3 — Lifelong Learning Record

### Definition

The Lifelong Learning Record is the comprehensive longitudinal record of an individual's complete learning journey across all educational phases and throughout their entire life.

### Lifecycle Coverage

```
PAUD (Early Childhood Education)
    ↓
SD (Primary School)
    ↓
SMP (Junior High School)
    ↓
SMA/SMK (Senior High School/Vocational School)
    ↓
Pendidikan Tinggi (Higher Education)
    ↓
Pelatihan (Professional Training)
    ↓
Sertifikasi (Certification)
    ↓
Dunia Kerja (Workplace Learning)
    ↓
Lifelong Learning (Continuous Upskilling/Reskilling)
```

### Record Components

- **Competencies**: All competencies developed across all phases
- **Achievements**: Academic and extracurricular achievements
- **Portfolios**: Complete portfolio of work and projects
- **Certifications**: Professional and skill certifications
- **Skills**: Technical, soft, and digital skills
- **Character Evidence**: Character development and values demonstration
- **Career Evidence**: Work experience and career progression

### Architectural Relationship

**Lifelong Learning Record is built from Learning Digital Twin.**

The Learning Digital Twin provides:
- The continuous data stream from educational phases
- The competency mastery progression over time
- The assessment evidence and portfolio artifacts
- The learning preferences and behavioral patterns

The Lifelong Learning Record adds:
- Cross-phase continuity and portability
- Longitudinal analysis and trend identification
- Career readiness and workforce planning data
- National human capital intelligence

---

## National Human Capital Intelligence Layer

### Intelligence Flow

```
Competency Graph
    ↓ structures
Learning Digital Twin
    ↓ aggregates
Lifelong Learning Record
    ↓ enables
National Human Capital Intelligence
```

### Strategic Foundation

This three-layer intelligence structure becomes the foundation for:

#### Personalization Learning
- Competency Graph provides the structure for personalized pathways
- Learning Digital Twin provides the context for individual adaptation
- Lifelong Learning Record provides the history for longitudinal personalization

#### Career Readiness
- Competency Graph defines required competencies for careers
- Learning Digital Twin tracks individual competency development
- Lifelong Learning Record demonstrates career readiness over time

#### Workforce Planning
- Competency Graph identifies national competency requirements
- Learning Digital Twin provides real-time competency data
- Lifelong Learning Record enables skills gap analysis and planning

#### Education Policy
- Competency Graph informs curriculum and standards development
- Learning Digital Twin provides evidence of learning effectiveness
- Lifelong Learning Record enables outcome-based policy evaluation

#### Indonesia Emas 2045
- Competency Graph aligns national education with human capital goals
- Learning Digital Twin enables measurement of progress toward goals
- Lifelong Learning Record provides data for strategic national planning

## Core Principle

**The National Learning Intelligence Backbone is the strategic foundation that enables NUSA to provide AI-native education at national scale while maintaining human-centered educational values.**

This backbone ensures that:
- AI recommendations are grounded in comprehensive competency understanding
- Personalization is based on complete learning context
- Lifelong learning is supported through continuous record-keeping
- National human capital development is data-driven and evidence-based
- Human judgment is supported by comprehensive intelligence

---

# SECTION 18 — AI Agent Coverage & 90% Automation Architecture

## Vision Statement

The Education Operating System (NUSA) envisions a future where:

**90% AI Assisted Operations**
**10% Human Judgment**

AI does not replace teachers. AI replaces administrative and repetitive work, enabling teachers to focus on high-value educational activities that require human expertise, empathy, and judgment.

## Human Authority Preservation

Humans maintain exclusive authority over:
- Empathy and emotional support
- Ethical judgment and decision-making
- Pedagogical decisions and teaching strategies
- Final assessment and grading
- Social-emotional interventions
- Complex disciplinary decisions
- Crisis management

AI assists but never replaces human educational judgment.

---

## AI Agent Landscape

| Domain | AI Agent | Responsibility |
| ------ | -------- | -------------- |
| **Curriculum** | Curriculum Agent | CP → TP transformation, TP → ATP generation, curriculum alignment verification |
| **Learning Planning** | ATP Agent | ATP generation, learning path optimization, resource matching |
| **Learning Planning** | Lesson Planning Agent | Modul Ajar drafting, lesson design support, activity sequencing |
| **Learning** | Learning Content Agent | Learning material generation, content adaptation, resource recommendation |
| **Assessment** | Assessment Agent | Rubric analysis, evidence classification, assessment item generation |
| **Reporting** | Reporting Agent | Narrative report generation, data synthesis, multi-language translation |
| **Student Wellbeing** | Student Success Agent | Engagement monitoring, early warning, intervention recommendation |
| **Parent Partnership** | Parent Communication Agent | Communication automation, translation, engagement analytics |
| **Teacher Growth** | Teacher Coach Agent | Professional growth recommendation, PD resource matching, classroom analysis |
| **School Improvement** | School Improvement Agent | School quality analytics, benchmarking, improvement strategy recommendation |
| **Quality Assurance** | Quality Assurance Agent | Compliance monitoring, gap analysis, accreditation support |
| **Career Readiness** | Career Guidance Agent | Career pathway recommendation, skill gap analysis, industry trend analysis |
| **Lifelong Learning** | Lifelong Learning Agent | Competency tracking, career readiness assessment, learning record management |
| **Data** | Education Analytics Agent | Learning analytics, predictive modeling, pattern recognition, insight generation |
| **AI Orchestration** | AI Governance Agent | Agent coordination, workflow orchestration, human-in-the-loop oversight |

---

## Automation Coverage Matrix

| Workflow | AI Automation | Human Validation | Target Automation |
| -------- | ------------- | ---------------- | ------------------ |
| **CP → TP Generation** | Automated transformation based on CP | Teacher review and approval | 95% |
| **TP → ATP Generation** | Automated sequencing and resource matching | Teacher review and adjustment | 90% |
| **Modul Ajar Drafting** | Automated content generation | Teacher customization and approval | 85% |
| **Learning Material Generation** | Automated resource creation and adaptation | Teacher review and selection | 80% |
| **Formative Assessment Analysis** | Automated scoring and pattern recognition | Teacher interpretation and intervention | 90% |
| **Portfolio Analysis** | Automated evidence classification and tagging | Teacher review and assessment | 85% |
| **Report Narrative Generation** | Automated narrative writing and translation | Teacher review and approval | 95% |
| **Parent Communication Draft** | Automated communication generation | Teacher review and sending | 90% |
| **Teacher Reflection Summary** | Automated reflection synthesis | Teacher review and action | 85% |
| **School Quality Analysis** | Automated analytics and benchmarking | Principal review and action | 90% |

**Overall Target**: 80–98% AI Automation across all workflows, with human validation at critical decision points.

---

## Human Governance Principle

### Human-in-the-Loop Governance

All AI Agents must operate under strict Human-in-the-Loop Governance:

#### Human Authority Domains

Humans maintain exclusive authority as:
- **Pedagogical Authority**: Final decisions on teaching strategies, learning approaches, and educational interventions
- **Ethical Authority**: Decisions involving ethical considerations, fairness, equity, and student welfare
- **Assessment Authority**: Final grading, promotion decisions, graduation decisions, and disciplinary actions

#### AI Agent Role

AI Agents serve exclusively as:
- **Co-pilot**: Supporting human decision-making with data-driven insights
- **Accelerator**: Accelerating routine tasks and reducing administrative burden
- **Recommender**: Providing recommendations that humans review, approve, modify, or reject

AI Agents are never decision-makers. AI Agents never own educational decisions. AI Agents never replace human judgment.

### Governance Mechanisms

- **Validation Points**: Defined human validation points for all AI-generated outputs
- **Approval Workflows**: Structured approval workflows for high-stakes decisions
- **Audit Trails**: Complete audit trails of AI recommendations and human decisions
- **Override Capabilities**: Human ability to override any AI recommendation
- **Transparency**: Explainable AI recommendations with clear rationale
- **Accountability**: Clear human accountability for all final decisions

---

# SECTION 19 — Domain Architecture Final Statement

## Architectural Position

The document `01_EDUCATION_DOMAIN_MODEL.md` now serves as the **Master Domain Architecture of the Indonesian Education Operating System (NUSA)**.

## Architectural Roles

This document functions as:

### Master Domain Architecture
- Definitive reference for all education domains in the Indonesian education system
- Complete domain model covering all national education functions
- Foundation for all subsequent architectural work
- Single source of truth for domain boundaries and relationships

### Foundation Architecture
- Stable foundation that can support decades of evolution
- Independent of technology implementation details
- Aligned with national education policy and vision
- Designed for scalability to serve millions of students

### Single Source of Truth for Education Domains
- Unambiguous definition of domain boundaries
- Clear ownership and accountability structures
- Well-defined cross-domain relationships
- Governed evolution process for domain changes

### Reference Model for Capability Mapping
- Foundation for capability identification and definition
- Mapping between domains and operational capabilities
- Basis for business process architecture
- Reference for service design and implementation

### Reference Model for AI Agent Design
- Clear mapping between domains and AI agents
- Definition of AI agent responsibilities and boundaries
- Human validation points for each AI agent
- Governance framework for AI agent behavior

### Reference Model for Future NUSA Development
- Architectural guardrails for all future development
- Decision framework for feature prioritization
- Integration patterns for new capabilities
- Evolution path for system growth

## Architectural Integration

This document integrates five critical architectural components:

### Domain Architecture
- 23 domains organized across 5 strategic categories
- Complete coverage of national education functions
- Clear domain boundaries and ownership
- Well-defined cross-domain relationships

### AI Architecture
- Multi-agent system with 15+ specialized AI agents
- 90% AI assistance with 10% human judgment
- Human-in-the-loop governance framework
- Clear human authority preservation

### Competency Architecture
- National Competency Graph as educational brain
- Learning Digital Twin for personalization
- Lifelong Learning Record for longitudinal tracking
- National Human Capital Intelligence layer

### Governance Architecture
- Human authority at all critical decision points
- Four-level decision framework (AI Autonomous, AI Assisted, Human Approval Required, Human Only)
- Clear accountability structures
- Ethical and educational judgment preservation

### Lifelong Learning Architecture
- National Learning Identity spanning entire lifecycle
- Cross-phase continuity and portability
- Career readiness and workforce planning support
- Foundation for Indonesia Emas 2045 human capital development

## Strategic Alignment

This domain architecture is fully aligned with:

### 00A_NATIONAL_EDUCATION_DIRECTION_2045.md
- Indonesia Emas 2045 vision and goals
- National human capital development objectives
- Graduate Profile (8 Dimensions) as North Star
- Outcome-driven education paradigm

### 00B_PRODUCT_VISION.md
- NUSA paradigm and principles
- AI-First Architecture vision
- 90% AI Assistance, 10% Human Judgment target
- Curriculum First, Pedagogy Before Technology principles

### 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md
- Platform strategy and architectural principles
- Human Authority Principle
- AI-Native Education Platform paradigm
- Governance and accountability frameworks

## Foundation for Subsequent Architecture Work

This domain model serves as the foundation for:

### 02_CAPABILITY_MODEL.md
- Capability identification based on domain structure
- Capability mapping to domains and AI agents
- Capability prioritization and dependency analysis
- Capability implementation roadmap

### 03_BUSINESS_PROCESS_ARCHITECTURE.md
- Business process design based on domain boundaries
- Cross-domain process orchestration
- Process automation opportunities
- Human-in-the-loop process design

### 04_AI_AGENT_ARCHITECTURE.md
- AI agent design based on domain responsibilities
- AI agent coordination and orchestration
- AI agent governance and oversight
- AI agent performance monitoring

### 05_DATA_ARCHITECTURE.md
- Data model design based on domain boundaries
- Single Source of Truth implementation
- Data flow and integration patterns
- Data governance and quality management

### 06_APPLICATION_ARCHITECTURE.md
- Application design based on domain structure
- Service boundaries and interfaces
- Integration patterns and APIs
- Scalability and performance considerations

## Final Architectural Mandate

This domain architecture establishes that:

- All domains, AI architecture, competency architecture, governance architecture, and lifelong learning architecture are fully integrated
- The architecture supports the vision of Indonesia Emas 2045
- The architecture enables AI-native education at national scale
- The architecture maintains human-centered educational values
- The architecture provides a stable foundation for decades of evolution

**This document is the Master Domain Architecture of the Indonesian Education Operating System (NUSA).**

---

# SECTION 20 — Capability Ownership Matrix

## Purpose

This section explicitly defines which domain owns each capability in the Education Operating System (NUSA). Clear capability ownership is essential for governance, roadmap planning, AI agent assignment, and establishing Single Source of Truth.

## Ownership Principles

### One Capability → One Primary Owner
- Each capability has exactly one primary domain owner
- Ownership determines governance, roadmap, AI agent assignment, and source of truth
- Primary owner has final authority over capability evolution

### Capability Ownership Not Shared
- Capabilities cannot have two primary owners
- Shared ownership creates ambiguity and accountability gaps
- Clear ownership enables effective decision-making

### Collaboration Allowed
- Other domains can collaborate on capability implementation
- Supporting domains provide inputs and consume outputs
- Collaboration does not transfer ownership

### Ownership Determines
- **Governance**: Primary owner defines capability standards and rules
- **Roadmap**: Primary owner plans capability evolution
- **AI Agent**: Primary owner assigns AI agent for capability
- **Source of Truth**: Primary owner maintains authoritative data

## Capability Ownership Matrix

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |

### Graduate Profile Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Graduate Profile Management | Graduate Profile Domain | All domains (alignment) |
| Graduate Outcome Definition | Graduate Profile Domain | Curriculum Domain |

### Curriculum Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| CP Management | Curriculum Domain | Graduate Profile Domain |
| TP Management | Curriculum Domain | Learning Planning Domain |
| ATP Management | Curriculum Domain | Learning Planning Domain |
| Curriculum Mapping | Curriculum Domain | Assessment Domain |

### Learning Planning Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Lesson Planning | Learning Planning Domain | Curriculum Domain |
| Modul Ajar Management | Learning Planning Domain | Learning Delivery Domain |

### Learning Delivery Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Learning Session | Learning Delivery Domain | Learning Planning Domain |
| Learning Activity | Learning Delivery Domain | Assessment Domain |

### Assessment Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Assessment Design | Assessment Domain | Curriculum Domain |
| Assessment Evidence | Assessment Domain | Learning Delivery Domain |
| Competency Evaluation | Assessment Domain | Graduate Profile Domain |

### Reporting Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Report Generation | Reporting Domain | Assessment Domain |
| Parent Reporting | Reporting Domain | Parent Partnership Domain |

### Student Development Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Student Growth Tracking | Student Wellbeing Domain | Assessment Domain |
| Intervention Planning | Student Wellbeing Domain | Teacher Growth Domain |

### Teacher Growth Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Teacher Reflection | Teacher Professional Growth Domain | School Improvement Domain |
| Professional Growth | Teacher Professional Growth Domain | School Improvement Domain |

### School Improvement Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| School Quality Planning | School Improvement Domain | Quality Assurance Domain |
| Improvement Monitoring | School Improvement Domain | Education Data Domain |

### Lifelong Learning Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Learning Passport | Lifelong Learning Record Domain | Career Domain |
| Credential Management | Lifelong Learning Record Domain | Quality Assurance Domain |

### Education Analytics Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| Learning Analytics | Education Data Domain | All domains (data sources) |
| Education Intelligence | Education Data Domain | AI Orchestration Domain |

### AI Governance Capabilities

| Capability | Primary Domain Owner | Supporting Domains |
| ---------- | -------------------- | ------------------ |
| AI Policy Management | AI Orchestration Domain | All domains (policy compliance) |
| AI Risk Monitoring | AI Orchestration Domain | Quality Assurance Domain |

## Strategic Implications

### For Capability Model
- Capability ownership provides foundation for capability identification
- Ownership boundaries define capability granularity
- Supporting domains identify cross-capability dependencies

### For Service Boundary
- Capability ownership defines service boundaries
- Primary owner determines service interface
- Supporting domains define service consumers

### For Team Boundary
- Capability ownership determines team responsibility
- Primary owner domain defines primary team
- Supporting domains define cross-team collaboration

### For AI Agent Assignment
- Each capability has assigned AI agent from primary domain
- AI agent responsibility aligns with capability ownership
- Human validation points defined by primary domain

## Core Principle

**Clear capability ownership is the foundation of effective governance and architectural coherence.**

This principle ensures that:
- Every capability has unambiguous ownership
- Governance structures are clear and effective
- AI agents have well-defined responsibilities
- System evolution is controlled and coherent

---

# SECTION 21 — Data Ownership Architecture

## Purpose

This section establishes the Single Source of Truth (SSOT) for all educational data entities in the Education Operating System (NUSA). Clear data ownership is essential for data consistency, governance, and system integrity.

## Data Ownership Principles

### Every Data Has One Owner
- Each data entity has exactly one domain owner
- Data owner has authority over data definition, quality, and evolution
- Data ownership is not shared to avoid ambiguity

### Every Data Has One Source of Truth
- Each data entity has one authoritative source
- Source of truth is maintained by the owning domain
- All other domains consume from the source of truth

### Data Sharing Allowed
- Data can be shared across domains for legitimate purposes
- Data sharing follows defined integration patterns
- Data consumers respect data ownership and governance

### Data Ownership Not Shared
- Data authority cannot be shared between domains
- Only one domain can make authoritative changes to data
- Data governance follows domain ownership

## Data Ownership Matrix

| Data Entity | Aggregate Root | Owning Domain | Source of Truth |
| ----------- | -------------- | ------------- | --------------- |
| **Graduate Profile** | Graduate Profile | Graduate Profile Domain | Graduate Profile Domain |
| **CP** | Curriculum | Curriculum Domain | Curriculum Domain |
| **TP** | Curriculum | Curriculum Domain | Curriculum Domain |
| **ATP** | Curriculum | Curriculum Domain | Curriculum Domain |
| **Modul Ajar** | Lesson Plan | Learning Planning Domain | Learning Planning Domain |
| **Learning Activity** | Learning Session | Learning Delivery Domain | Learning Delivery Domain |
| **Assessment Evidence** | Assessment | Assessment Domain | Assessment Domain |
| **Student Profile** | Student | Student Wellbeing Domain | Student Wellbeing Domain |
| **Competency Record** | Competency Record | Lifelong Learning Record Domain | Lifelong Learning Record Domain |
| **Learning Portfolio** | Portfolio | Lifelong Learning Record Domain | Lifelong Learning Record Domain |
| **Teacher Profile** | Teacher | Teacher Professional Growth Domain | Teacher Professional Growth Domain |
| **School Improvement Plan** | Improvement Plan | School Improvement Domain | School Improvement Domain |
| **School Quality Data** | School Quality | Quality Assurance Domain | Quality Assurance Domain |
| **Career Profile** | Career | Career & Future Readiness Domain | Career & Future Readiness Domain |
| **Learning Analytics** | Analytics | Education Data Domain | Education Data Domain |
| **AI Agent Logs** | AI Agent | AI Orchestration Domain | AI Orchestration Domain |

## Core Principles

### No Duplicate Authority
- Data authority cannot be duplicated across domains
- Only one domain has authority to make authoritative changes
- This prevents data inconsistency and governance conflicts

### Data Replication Allowed
- Data can be replicated for performance, availability, or locality
- Replicated data is clearly marked as non-authoritative
- Replicated data must sync with source of truth

### Authority Not Replicated
- Data authority cannot be replicated
- Only the owning domain has authority
- Replicated data does not confer authority

### Data Governance Follows Ownership
- Data governance is the responsibility of the owning domain
- Data quality standards are defined by the owning domain
- Data access policies are controlled by the owning domain

## Strategic Implications

### For Data Architecture
- Clear data ownership enables proper data model design
- Source of truth patterns prevent data duplication
- Data integration patterns follow ownership boundaries

### For System Architecture
- Service boundaries align with data ownership
- API design respects data authority
- Data flow patterns follow ownership hierarchy

### For Governance
- Data governance is clear and effective
- Data quality accountability is unambiguous
- Data access control is properly defined

### For AI
- AI agents respect data ownership
- AI recommendations are based on authoritative data
- AI-generated data follows ownership rules

## Core Principle

**Every data entity has one owner, one source of truth, and unambiguous governance.**

This principle ensures that:
- Data consistency is maintained across the system
- Data governance is clear and effective
- System integrity is preserved
- AI operations are based on authoritative data

---

# SECTION 22 — Domain Dependency Architecture

## Purpose

This section defines the official dependency hierarchy among domains in the Education Operating System (NUSA). Understanding domain dependencies is critical for implementation sequencing, system integration, and architectural governance.

## Dependency Hierarchy

### Level 1 — Strategic Domains (Core Education Flow)

The foundational domains that define the core educational process:

```
Graduate Profile Domain
    ↓ defines
Curriculum Domain
    ↓ enables
Learning Planning Domain
    ↓ enables
Learning Delivery Domain
    ↓ generates
Assessment Domain
    ↓ produces
Reporting Domain
```

**Dependency Rule**: If a domain in this hierarchy is unavailable, all downstream domains cannot function fully.

**Explanation**:
- Graduate Profile defines desired outcomes
- Curriculum translates outcomes into learning capabilities (CP)
- Learning Planning translates CP into actionable plans (TP, ATP)
- Learning Delivery executes the plans
- Assessment collects evidence of learning
- Reporting communicates progress

---

### Level 2 — Growth Domains (Quality Improvement)

Domains that depend on core educational data to drive improvement:

```
Assessment Domain
    ↓ provides data for
Student Development Domain
    ↓ informs
Teacher Growth Domain
    ↓ enables
School Improvement Domain
```

**Dependency Rule**: Growth domains require data from core domains to function effectively.

**Explanation**:
- Assessment provides evidence of student learning
- Student Development uses assessment data for growth tracking
- Teacher Growth uses student and assessment data for professional development
- School Improvement uses all data for quality planning

---

### Level 3 — Intelligence Domains (Analytics and AI)

Domains that aggregate and analyze data to provide intelligence:

```
Assessment Domain
    ↓ provides data to
Education Analytics Domain
    ↓ enables
AI Governance Domain
    ↓ powers
National Intelligence Layer
```

**Dependency Rule**: Intelligence domains require data from operational domains to provide value.

**Explanation**:
- Assessment and other domains provide raw data
- Education Analytics processes and analyzes data
- AI Governance ensures AI operates within defined parameters
- National Intelligence Layer provides strategic insights

---

### Level 4 — Lifelong Learning Domains (Longitudinal Tracking)

Domains that track learning across the entire lifecycle:

```
Assessment Domain
    ↓ generates
Competency Record
    ↓ contributes to
Learning Portfolio
    ↓ feeds
Digital Twin
    ↓ aggregates into
Lifelong Learning Record
    ↓ enables
Career Readiness
    ↓ supports
Workforce Readiness
```

**Dependency Rule**: Lifelong learning domains require continuous data flow from operational domains.

**Explanation**:
- Assessment provides competency evidence
- Competency Record tracks mastery over time
- Learning Portfolio stores artifacts and achievements
- Digital Twin provides comprehensive learning representation
- Lifelong Learning Record maintains longitudinal history
- Career Readiness uses learning record for guidance
- Workforce Readiness aggregates for national planning

---

## Dependency Principles

### Upstream Dependency
- Downstream domains depend on upstream domains
- Upstream domains can function independently
- Downstream domains cannot function without upstream

### Data Flow Direction
- Data flows from upstream to downstream
- Downstream domains consume upstream data
- Upstream domains are not dependent on downstream

### Implementation Sequencing
- Upstream domains should be implemented before downstream
- This ensures data availability for dependent domains
- Phased implementation follows dependency hierarchy

### Failure Impact
- Failure in upstream domain affects all downstream domains
- Failure in downstream domain does not affect upstream domains
- Critical dependencies must be highly available

## Strategic Implications

### For Implementation
- Implementation roadmap must follow dependency hierarchy
- MVP should focus on Level 1 (Strategic Domains)
- Subsequent phases add Level 2, 3, and 4

### For System Architecture
- System design must respect dependency boundaries
- Integration patterns follow dependency hierarchy
- Data pipelines align with dependency flow

### For Governance
- Domain governance follows dependency structure
- Upstream domains have governance authority over data
- Downstream domains respect upstream governance

### For Risk Management
- Critical dependencies identified and monitored
- Redundancy planned for critical upstream domains
- Failure recovery prioritizes upstream domains

## Core Principle

**Domain dependencies define the official hierarchy of NUSA implementation and integration.**

This principle ensures that:
- Implementation follows logical dependency sequence
- System architecture respects domain boundaries
- Data flow aligns with dependency hierarchy
- Risk management focuses on critical dependencies

---

# SECTION 23 — Architectural Governance Rules

## Purpose

This section establishes the official architectural governance rules that all NUSA development must follow. These rules ensure architectural coherence, alignment with educational principles, and system integrity.

## Rule 1: Curriculum Before Feature

**Statement**: All features must be traceable to the curriculum hierarchy.

**Traceability Requirement**:
```
Graduate Profile (8 Dimensions)
    ↓
CP (Capaian Pembelajaran)
    ↓
TP (Tujuan Pembelajaran)
    ↓
ATP (Learning Path)
```

**Implementation**:
- Every feature request must demonstrate curriculum traceability
- Features without curriculum traceability are rejected
- Curriculum traceability is documented in feature specifications

**Rationale**: Ensures that all system features contribute to educational outcomes defined in the national curriculum.

---

## Rule 2: Evidence Before Reporting

**Statement**: All reports must be based on evidence.

**Evidence Requirement**:
- Reports must be derived from assessment evidence
- Reports without evidence backing are prohibited
- Evidence sources must be documented in reports

**Implementation**:
- Report generation requires evidence data
- Evidence validation before report publication
- Audit trail from evidence to report

**Rationale**: Ensures that all communications are based on actual student performance data, not assumptions or estimates.

---

## Rule 3: Human Before AI Decision

**Statement**: AI may recommend, but humans must decide.

**Decision Authority**:
- AI provides recommendations and analysis
- Humans make final educational decisions
- AI never has autonomous decision authority

**Implementation**:
- All AI outputs require human validation
- Critical decisions require explicit human approval
- AI recommendations are clearly marked as such

**Rationale**: Maintains human authority and accountability in education while leveraging AI for data-driven insights.

---

## Rule 4: Single Source of Truth

**Statement**: Each data entity has one domain owner.

**Ownership Requirement**:
- Every data entity has exactly one owning domain
- Data authority is not shared
- Source of truth is maintained by owning domain

**Implementation**:
- Data ownership documented in Data Ownership Matrix
- Data changes authorized only by owning domain
- Data replication clearly marked as non-authoritative

**Rationale**: Prevents data inconsistency and governance conflicts through clear ownership.

---

## Rule 5: Outcome Traceability

**Statement**: All learning activities must trace to Profil Lulusan 8 Dimensi.

**Traceability Chain**:
```
Learning Activity
    ↓ contributes to
Competency Achievement
    ↓ contributes to
Learning Outcome
    ↓ contributes to
Graduate Profile (8 Dimensions)
    ↓ contributes to
Indonesia Emas 2045
```

**Implementation**:
- Every activity tagged with target competencies
- Competency mapping to Graduate Profile dimensions
- Traceability audit trail maintained

**Rationale**: Ensures that all educational activities contribute to national human capital goals.

---

## Rule 6: AI Must Reduce Teacher Work

**Statement**: AI must not increase administrative burden on teachers.

**Burden Reduction Requirement**:
- AI features must reduce teacher workload
- AI features must not add administrative tasks
- Teacher time saved is a success metric

**Implementation**:
- Workload impact assessment for all AI features
- Teacher feedback on burden reduction
- Features that increase burden are rejected

**Rationale**: Ensures that AI serves its purpose of reducing teacher burden, not increasing it.

---

## Rule 7: Human-Centered Education

**Statement**: System purpose is to enhance human relationships, not replace them.

**Human-Centered Requirement**:
- System enhances teacher-student relationships
- System enhances parent-teacher communication
- System enhances teacher collaboration
- System does not replace human interaction

**Implementation**:
- Human relationship impact assessment
- Design focuses on human connection
- Technology supports, not replaces, humans

**Rationale**: Ensures that technology serves the fundamental human relationships that define quality education.

---

## Governance Enforcement

### Architectural Review Board
- All major features must pass architectural review
- Review checks compliance with governance rules
- Non-compliant features are rejected or require modification

### Automated Validation
- Automated checks for traceability requirements
- Automated validation of data ownership
- Automated monitoring of AI-human decision boundaries

### Audit Mechanisms
- Regular audits of system compliance with rules
- Audit reports reviewed by governance board
- Non-compliance triggers remediation actions

## Core Principle

**Architectural governance rules ensure that NUSA remains aligned with educational principles and human values.**

These rules ensure that:
- System features contribute to educational outcomes
- Human authority is maintained
- Data integrity is preserved
- Teacher burden is reduced
- Human relationships are enhanced

---

# SECTION 24 — Foundation Architecture Freeze v1.0

## Declaration

The following documents now form the **National Education Foundation Architecture** for the Indonesian Education Operating System (NUSA):

### Foundation Documents

1. **00A_NATIONAL_EDUCATION_DIRECTION_2045.md**
   - Vision Architecture
   - National education direction to 2045
   - Indonesia Emas 2045 goals
   - Graduate Profile (8 Dimensions) definition

2. **00B_PRODUCT_VISION.md**
   - Product Architecture
   - NUSA product vision and strategy
   - AI-First Architecture paradigm
   - 90% AI Assistance, 10% Human Judgment target

3. **00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md**
   - Operating Principles
   - Platform strategy and architectural principles
   - Human Authority Principle
   - AI-Native Education Platform paradigm

4. **01_EDUCATION_DOMAIN_MODEL.md**
   - Domain Architecture
   - Complete domain model with 23 domains
   - Capability ownership and data ownership
   - Domain dependencies and governance rules

## Foundation Architecture Components

The National Education Foundation Architecture consists of:

### Vision Architecture
- National education vision and goals
- Graduate Profile as North Star
- Outcome-driven education paradigm
- Indonesia Emas 2045 alignment

### Product Architecture
- NUSA product vision and strategy
- AI-First Architecture principles
- Platform layer architecture
- Single Source of Truth architecture

### Operating Principles
- Human Authority Principle
- Curriculum First principle
- Pedagogy Before Technology principle
- AI-Native Education Platform paradigm

### Domain Architecture
- 23 domains across 5 strategic categories
- Capability ownership matrix
- Data ownership architecture
- Domain dependency hierarchy
- Architectural governance rules

## Official Reference Status

This foundation architecture is the **sole official reference** for all subsequent design work:

### For Capability Architecture
- 02_CAPABILITY_MODEL.md must align with domain ownership
- Capability boundaries must respect domain boundaries
- Capability dependencies must follow domain dependencies

### For Business Process Architecture
- 03_BUSINESS_PROCESS_ARCHITECTURE.md must align with domain model
- Process boundaries must respect domain boundaries
- Process flows must follow dependency hierarchy

### For Data Architecture
- 04_DATA_ARCHITECTURE.md must align with data ownership
- Data models must respect aggregate root definitions
- Data flows must follow dependency hierarchy

### For AI Architecture
- 05_AI_ARCHITECTURE.md must align with AI agent mapping
- AI agents must respect domain ownership
- AI governance must follow architectural governance rules

### For Application Architecture
- 06_APPLICATION_ARCHITECTURE.md must align with domain boundaries
- Service boundaries must respect domain boundaries
- Integration patterns must follow dependency hierarchy

### For MVP Development
- MVP 30 Hari must follow domain prioritization
- MVP scope must respect dependency hierarchy
- MVP features must comply with governance rules

## Architectural Compliance Requirement

**No capability, feature, service, AI agent, workflow, or database entity may contradict this foundation architecture.**

### Compliance Checklist
- [ ] Aligned with 00A (National Education Direction)
- [ ] Aligned with 00B (Product Vision)
- [ ] Aligned with 00C (Operating Principles)
- [ ] Aligned with 01 (Domain Model)
- [ ] Respects capability ownership
- [ ] Respects data ownership
- [ ] Follows domain dependencies
- [ ] Complies with governance rules

### Non-Compliance Consequence
Any design that contradicts this foundation architecture must be rejected or modified to achieve compliance.

## Freeze Status

This foundation architecture is declared **Frozen v1.0** as of June 2026.

### What Frozen Means
- Foundation documents are stable and authoritative
- Changes require formal architectural review process
- All subsequent work must comply with frozen foundation
- Foundation evolution follows governed process

### What Can Still Evolve
- Domain implementations can evolve within boundaries
- Capabilities can be added within domain ownership
- AI agents can be enhanced within governance rules
- Technology choices can change within architectural constraints

## Strategic Mandate

The National Education Foundation Architecture v1.0 is established as:

**The authoritative foundation for all Education Operating System (NUSA) development.**

This mandate ensures that:
- All development is aligned with national education vision
- All architecture is coherent and consistent
- All systems are built on stable foundations
- All evolution is governed and controlled

## Final Statement

**The Foundation Architecture Phase is complete.**

The foundation documents (00A, 00B, 00C, 01) provide a complete, coherent, and authoritative foundation for the Education Operating System Indonesia 2045.

With this foundation frozen, the risk for subsequent architecture work (02_CAPABILITY_MODEL.md, 03_BUSINESS_PROCESS_ARCHITECTURE.md, 04_DATA_ARCHITECTURE.md, 05_AI_ARCHITECTURE.md, 06_APPLICATION_ARCHITECTURE.md, MVP 30 Hari) is significantly reduced.

**Foundation Architecture Freeze v1.0 is declared.**

---

# SECTION 25 — Domain Prioritization
22. **Lifelong Learning Record** - National learning identity, longitudinal record
23. **AI Orchestration Domain** - Platform layer, agent coordination

**Success Criteria**: Workload optimized, accreditation supported, career readiness enabled, lifelong learning tracked, AI orchestration operational

---

# SECTION 26 — Strategic Conclusion

## What Should Be Built First

**Curriculum Operating System Core**:
1. Graduate Profile Framework (8 Dimensions)
2. Curriculum (CP) Repository
3. Learning Planning Tools (TP, ATP, Modul Ajar with AI)
4. Assessment Tools (formative/summative with AI)
5. Reporting Tools (automated with AI)
6. AI Orchestration Module (multi-agent coordination)

**Why This Core First**:
- Immediate burden reduction for teachers
- Foundational to all other educational work
- Direct policy alignment with Permendikdasmen 2025
- Teacher adoption through visible value
- AI-Native architecture from day one

## What Should NOT Be Built First

**Traditional SIS Features**:
- Complex attendance tracking
- Timetable/scheduling optimization
- Financial management
- Inventory management
- Complex dashboards for administrators

**Why Not These First**:
- Do not directly reduce core teacher burden
- Can be addressed with basic tools initially
- Do not support the core educational mission
- Can distract from curriculum operating system focus
- Violate the Curriculum First principle

## Why Education Operating System > Traditional SIS

**Traditional SIS Focus**:
- Administrative efficiency
- Data management
- Reporting compliance
- School operations

**Education Operating System Focus**:
- Educational effectiveness
- Teacher capacity
- Student learning
- Continuous improvement
- AI-Native experience

**Key Difference**:
- SIS manages the school as an organization
- NUSA manages the core work of education
- SIS is necessary but not sufficient
- NUSA is transformative

## Strategic Position

- Indonesia is implementing Permendikdasmen 2025 nationwide
- Teachers need support to implement Profil Lulusan 8 Dimensi
- Administrative burden prevents effective implementation
- NUSA directly addresses this challenge
- Success depends on teacher capacity, not administrative efficiency
- AI-Native architecture enables scale to millions of students

## Implementation Strategy

1. Build curriculum core first (Graduate Profile, CP, Planning, Assessment, Reporting)
2. Implement AI Orchestration Module from day one
3. Demonstrate value through teacher burden reduction
4. Build trust through educational effectiveness
5. Expand to support all educational domains
6. Add administrative features as needed, not as priority

## Success Metrics

- Teacher time saved on planning, assessment, reporting (target: 90% reduction)
- Teacher satisfaction with tools
- Alignment with Permendikdasmen 2025 requirements
- Student learning outcomes improvement (target: graduate profile development)
- School quality improvement
- AI adoption rate (target: 90% of workflows AI-assisted)

## Foundation Lock

This domain model is locked and serves as the official reference for:
- Domain-Driven Design (DDD)
- Event Storming
- Bounded Context
- Database Architecture
- AI Agent Architecture
- Product Roadmap
- MVP Development

All subsequent development must respect the domain relationships and dependencies mapped here. Software architecture should follow domain architecture, not the other way around.

## Final Principle

> Every feature, every workflow, every AI agent, and every data model must ultimately contribute to the development of Indonesia's Graduate Profile (8 Dimensions) and the realization of Indonesia Emas 2045.

This is the ultimate test for every architectural decision, every feature request, every system design, and every implementation. If it does not contribute to the development of Profil Lulusan 8 Dimensi and the realization of Indonesia Emas 2045, it should not exist in NUSA.

---

# SECTION 27 — Document Status

**Version**: 8.0
**Status**: FOUNDATION DOCUMENT - LOCKED
**Alignment**: 100% aligned with 00A_NATIONAL_EDUCATION_DIRECTION_2045.md, 00B_PRODUCT_VISION.md, and 00C_EDUCATION_OPERATING_SYSTEM_PRINCIPLES.md
**Purpose**: Domain model for Education Operating System Indonesia 2045
**Governance**: Changes to this document require review and approval from Chief Education Architect

This document is locked and serves as the official reference for:
- Domain-Driven Design (DDD)
- Event Storming
- Bounded Context
- Database Architecture
- AI Agent Architecture
- Product Roadmap
- MVP Development

## Version History

- **Version 8.0** (June 2026): Added Capability Ownership Matrix (explicit capability ownership across all domains), Data Ownership Architecture (Single Source of Truth for all data entities), Domain Dependency Architecture (official dependency hierarchy across 4 levels), Architectural Governance Rules (7 official governance rules), and Foundation Architecture Freeze v1.0 (declaration of National Education Foundation Architecture with 00A, 00B, 00C, 01). Document now serves as frozen Foundation Architecture v1.0, ready as official reference for 02_CAPABILITY_MODEL.md, 03_BUSINESS_PROCESS_ARCHITECTURE.md, 04_DATA_ARCHITECTURE.md, 05_AI_ARCHITECTURE.md, 06_APPLICATION_ARCHITECTURE.md, and MVP 30 Hari with significantly reduced architectural risk.
- **Version 7.0** (June 2026): Added National Education Domain Coverage (comprehensive coverage matrix of all national education functions), National Learning Intelligence Backbone (three-layer intelligence architecture: Competency Graph → Learning Digital Twin → Lifelong Learning Record → National Human Capital Intelligence), AI Agent Coverage & 90% Automation Architecture (AI Agent Landscape, Automation Coverage Matrix, Human Governance Principle), and Domain Architecture Final Statement. Document now serves as Master Domain Architecture of Indonesian Education Operating System (NUSA), fully aligned with 00A, 00B, and 00C, ready as foundation for 02_CAPABILITY_MODEL.md, 03_BUSINESS_PROCESS_ARCHITECTURE.md, 04_AI_AGENT_ARCHITECTURE.md, 05_DATA_ARCHITECTURE.md, and 06_APPLICATION_ARCHITECTURE.md.
- **Version 6.0** (June 2026): Added Education Capability Architecture (Domain → Capability → AI separation), Lifelong Learning Record Architecture (National Learning Identity), AI Orchestration Architecture (Multi-Agent System), Human Governance Layer (90% AI / 10% Human decision levels), and Outcome Traceability Matrix. Document now serves as complete Enterprise Education Business Architecture ready for Capability Model, Business Process Architecture, AI Agent Architecture, Data Architecture, and Application Architecture.
- **Version 5.0** (June 2026): Added Competency Graph Architecture and Learning Digital Twin Architecture sections. Updated Lifelong Learning Record classification from Foundation to Strategic domain. Updated domain classification counts (Core: 6, Strategic: 6, Supporting: 6, Foundation: 3, Platform: 2).
- **Version 4.0** (June 2026): Added Education Outcome Hierarchy, Core vs Supporting Domain Classification, Bounded Context Map, Aggregate Root Definitions, Domain Events Architecture, AI Agent ↔ Domain Mapping, and Lifelong Learning Record Domain. Expanded from 22 to 23 domains.
- **Version 3.0** (June 2026): Initial major revision aligning with NUSA paradigm, AI-First Architecture, and Permendikdasmen 2025.
- **Version 2.0**: Previous version with 15 domains.

---

**Curriculum Before Code. Pedagogy Before Technology. AI-First Architecture. Human Flourishing Through AI.**
