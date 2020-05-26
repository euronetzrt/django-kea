import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="django-kea",
    version='0.0.3',
    author="Richard Kojedzinszky",
    author_email="richard@kojedz.in",
    description="Django models and admin for ISC Kea dhcp server",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/euronetzrt/django-kea",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
)
